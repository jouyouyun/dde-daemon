/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package network

import (
	"pkg.linuxdeepin.com/lib/dbus"
	"sync"
)

const (
	dbusNetworkDest = "com.deepin.daemon.Network"
	dbusNetworkPath = "/com/deepin/daemon/Network"
	dbusNetworkIfs  = "com.deepin.daemon.Network"
)

const (
	opAdded = iota
	opRemoved
	opUpdated
)

type connectionData map[string]map[string]dbus.Variant

type Manager struct {
	config *config

	// update by manager.go
	State uint32 // networking state

	activeConnectionsLocker sync.Mutex
	activeConnections       []*activeConnection
	ActiveConnections       string // array of connections that activated and marshaled by json

	NetworkingEnabled bool `access:"readwrite"` // airplane mode for NetworkManager
	WirelessEnabled   bool `access:"readwrite"`
	WwanEnabled       bool `access:"readwrite"`
	WiredEnabled      bool `access:"readwrite"`
	VpnEnabled        bool `access:"readwrite"`

	// update by manager_devices.go
	devicesLocker sync.Mutex
	devices       map[string][]*device
	Devices       string // array of device objects and marshaled by json

	accessPointsLocker sync.Mutex
	accessPoints       map[dbus.ObjectPath][]*accessPoint

	// update by manager_connections.go
	connectionsLocker sync.Mutex
	connections       map[string][]*connection
	Connections       string // array of connection information and marshaled by json

	connectionSessionsLocker sync.Mutex
	connectionSessions       []*ConnectionSession

	// signals
	NeedSecrets                  func(string, string, string)
	DeviceStateChanged           func(devPath string, newState uint32) // TODO remove
	AccessPointAdded             func(devPath, apJSON string)
	AccessPointRemoved           func(devPath, apJSON string)
	AccessPointPropertiesChanged func(devPath, apJSON string)
	DeviceEnabled                func(devPath string, enabled bool)

	agent         *agent
	stateNotifier *stateNotifier
}

func (m *Manager) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		Dest:       dbusNetworkDest,
		ObjectPath: dbusNetworkPath,
		Interface:  dbusNetworkIfs,
	}
}

// initialize slice code
func initSlices() {
	initNmDbusObjects()
	initProxyGsettings()
	initAvailableValuesSecretFlags()
	initAvailableValuesNmPptpSecretFlags()
	initAvailableValuesNmL2tpSecretFlags()
	initAvailableValuesNmVpncSecretFlags()
	initAvailableValuesNmOpenvpnSecretFlags()
	initAvailableValuesWirelessChannel()
	initAvailableValues8021x()
	initAvailableValuesIp4()
	initAvailableValuesIp6()
	initNmStateReasons()
}

func NewManager() (m *Manager) {
	m = &Manager{}
	m.config = newConfig()
	return
}

func DestroyManager(m *Manager) {
	destroyStateNotifier(m.stateNotifier)
	destroyAgent(m.agent)
	m.clearConnectionSessions()
	dbus.UnInstallObject(m)
}

func (m *Manager) initManager() {
	// setup global switches
	m.initPropNetworkingEnabled()
	nmManager.NetworkingEnabled.ConnectChanged(func() {
		m.setPropNetworkingEnabled()
	})
	m.initPropWirelessEnabled()
	nmManager.WirelessEnabled.ConnectChanged(func() {
		m.setPropWirelessEnabled()
	})
	m.initPropWwanEnabled()
	nmManager.WwanEnabled.ConnectChanged(func() {
		m.setPropWwanEnabled()
	})

	// load virtual global switches information from configuration file
	m.initPropWiredEnabled()
	m.initPropVpnEnabled()

	// initialize device and connection managers
	m.initDeviceManage()
	m.initConnectionManage()

	// update property "ActiveConnections" after devices initialized
	nmManager.ActiveConnections.ConnectChanged(func() {
		m.updateActiveConnections()
	})
	m.updateActiveConnections()

	// update property "State"
	nmManager.State.ConnectChanged(func() {
		m.setPropState()
	})
	m.setPropState()

	m.agent = newAgent()
	m.stateNotifier = newStateNotifier()
}

func (m *Manager) updateActiveConnections() {
	m.activeConnectionsLocker.Lock()
	defer m.activeConnectionsLocker.Unlock()

	// reset all exists active connection objects
	for i, _ := range m.activeConnections {
		// destroy object to reset all property connects
		if m.activeConnections[i].nmVpnConn != nil {
			nmDestroyVpnConnection(m.activeConnections[i].nmVpnConn)
		}
		nmDestroyActiveConnection(m.activeConnections[i].nmAConn)
		m.activeConnections[i] = nil
	}
	m.activeConnections = make([]*activeConnection, 0)
	for _, acPath := range nmGetActiveConnections() {
		if nmAConn, err := nmNewActiveConnection(acPath); err == nil {
			aconn := &activeConnection{
				nmAConn: nmAConn,
				path:    acPath,
				Devices: nmAConn.Devices.Get(),
				Uuid:    nmAConn.Uuid.Get(),
				State:   nmAConn.State.Get(),
				Vpn:     nmAConn.Vpn.Get(),
			}
			if cpath, err := nmGetConnectionByUuid(aconn.Uuid); err == nil {
				aconn.Id = nmGetConnectionId(cpath)
			}

			nmAConn.State.ConnectChanged(func() {
				m.activeConnectionsLocker.Lock()
				defer m.activeConnectionsLocker.Unlock()

				aconn.State = nmAConn.State.Get()
				m.setPropActiveConnections()
			})
			aconn.State = nmAConn.State.Get()

			// dispatch vpn connection
			if aconn.Vpn {
				if nmVpnConn, err := nmNewVpnConnection(acPath); err == nil {
					aconn.nmVpnConn = nmVpnConn
					nmVpnConn.ConnectVpnStateChanged(func(state, reason uint32) {
						// update vpn config
						m.config.setVpnConnectionActivated(aconn.Uuid, isVpnConnectionStateInActivating(state))

						// notification for vpn
						if isVpnConnectionStateActivated(state) {
							notifyVpnConnected(aconn.Id)
						} else if isVpnConnectionStateDeactivate(state) {
							notifyVpnDisconnected(aconn.Id)
						} else if isVpnConnectionStateFailed(state) {
							notifyVpnFailed(aconn.Id, reason)
						}
					})
				}
			}

			m.activeConnections = append(m.activeConnections, aconn)
		}
	}
	m.setPropActiveConnections()
}
