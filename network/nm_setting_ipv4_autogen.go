// This file is automatically generated, please don't edit manully."
package main

import (
	"fmt"
)

// Get key type
func getSettingIp4ConfigKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_IP4_CONFIG_METHOD:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_DNS:
		t = ktypeWrapperIpv4Dns
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		t = ktypeWrapperIpv4Addresses
	case NM_SETTING_IP4_CONFIG_ROUTES:
		t = ktypeWrapperIpv4Routes
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		t = ktypeBoolean
	}
	return
}

// Check is key in current setting field
func isKeyInSettingIp4Config(key string) bool {
	switch key {
	case NM_SETTING_IP4_CONFIG_METHOD:
		return true
	case NM_SETTING_IP4_CONFIG_DNS:
		return true
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		return true
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		return true
	case NM_SETTING_IP4_CONFIG_ROUTES:
		return true
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		return true
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		return true
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		return true
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		return true
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		return true
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		return true
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		return true
	}
	return false
}

// Get key's default value
func getSettingIp4ConfigKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		Logger.Error("invalid key:", key)
	case NM_SETTING_IP4_CONFIG_METHOD:
		valueJSON = `""`
	case NM_SETTING_IP4_CONFIG_DNS:
		valueJSON = `null`
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		valueJSON = `""`
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		valueJSON = `null`
	case NM_SETTING_IP4_CONFIG_ROUTES:
		valueJSON = `null`
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		valueJSON = `false`
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		valueJSON = `false`
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		valueJSON = `""`
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		valueJSON = `true`
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		valueJSON = `""`
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		valueJSON = `false`
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		valueJSON = `true`
	}
	return
}

// Get JSON value generally
func generalGetSettingIp4ConfigKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		Logger.Error("generalGetSettingIp4ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP4_CONFIG_METHOD:
		value = getSettingIp4ConfigMethodJSON(data)
	case NM_SETTING_IP4_CONFIG_DNS:
		value = getSettingIp4ConfigDnsJSON(data)
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		value = getSettingIp4ConfigDnsSearchJSON(data)
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		value = getSettingIp4ConfigAddressesJSON(data)
	case NM_SETTING_IP4_CONFIG_ROUTES:
		value = getSettingIp4ConfigRoutesJSON(data)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		value = getSettingIp4ConfigIgnoreAutoRoutesJSON(data)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		value = getSettingIp4ConfigIgnoreAutoDnsJSON(data)
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		value = getSettingIp4ConfigDhcpClientIdJSON(data)
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		value = getSettingIp4ConfigDhcpSendHostnameJSON(data)
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		value = getSettingIp4ConfigDhcpHostnameJSON(data)
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		value = getSettingIp4ConfigNeverDefaultJSON(data)
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		value = getSettingIp4ConfigMayFailJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingIp4ConfigKeyJSON(data _ConnectionData, key, valueJSON string) {
	switch key {
	default:
		Logger.Error("generalSetSettingIp4ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP4_CONFIG_METHOD:
		logicSetSettingIp4ConfigMethodJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DNS:
		setSettingIp4ConfigDnsJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		setSettingIp4ConfigDnsSearchJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		setSettingIp4ConfigAddressesJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_ROUTES:
		setSettingIp4ConfigRoutesJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		setSettingIp4ConfigIgnoreAutoRoutesJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		setSettingIp4ConfigIgnoreAutoDnsJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		setSettingIp4ConfigDhcpClientIdJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		setSettingIp4ConfigDhcpSendHostnameJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		setSettingIp4ConfigDhcpHostnameJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		setSettingIp4ConfigNeverDefaultJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		setSettingIp4ConfigMayFailJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingIp4ConfigMethodExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD)
}
func isSettingIp4ConfigDnsExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS)
}
func isSettingIp4ConfigDnsSearchExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH)
}
func isSettingIp4ConfigAddressesExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES)
}
func isSettingIp4ConfigRoutesExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES)
}
func isSettingIp4ConfigIgnoreAutoRoutesExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES)
}
func isSettingIp4ConfigIgnoreAutoDnsExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS)
}
func isSettingIp4ConfigDhcpClientIdExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID)
}
func isSettingIp4ConfigDhcpSendHostnameExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME)
}
func isSettingIp4ConfigDhcpHostnameExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME)
}
func isSettingIp4ConfigNeverDefaultExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT)
}
func isSettingIp4ConfigMayFailExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL)
}

// Ensure field and key exists and not empty
func ensureFieldSettingIp4ConfigExists(data _ConnectionData, errs map[string]string, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME) {
		rememberError(errs, relatedKey, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_IP4_CONFIG_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_IP4_CONFIG_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_IP4_CONFIG_SETTING_NAME))
	}
}
func ensureSettingIp4ConfigMethodNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigMethodExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_METHOD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigMethod(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_METHOD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigDnsNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigDnsExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDns(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DNS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigDnsSearchNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigDnsSearchExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DNS_SEARCH, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDnsSearch(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DNS_SEARCH, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigAddressesNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigAddressesExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_ADDRESSES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigAddresses(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_ADDRESSES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigRoutesNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigRoutes(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_ROUTES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigIgnoreAutoRoutesNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigIgnoreAutoRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigIgnoreAutoDnsNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigIgnoreAutoDnsExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigDhcpClientIdNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigDhcpClientIdExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDhcpClientId(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigDhcpSendHostnameNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigDhcpSendHostnameExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigDhcpHostnameNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigDhcpHostnameExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDhcpHostname(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigNeverDefaultNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigNeverDefaultExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigMayFailNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp4ConfigMayFailExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_MAY_FAIL, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingIp4ConfigMethod(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD).(string)
	return
}
func getSettingIp4ConfigDns(data _ConnectionData) (value []uint32) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS).([]uint32)
	return
}
func getSettingIp4ConfigDnsSearch(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH).(string)
	return
}
func getSettingIp4ConfigAddresses(data _ConnectionData) (value [][]uint32) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES).([][]uint32)
	return
}
func getSettingIp4ConfigRoutes(data _ConnectionData) (value [][]uint32) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES).([][]uint32)
	return
}
func getSettingIp4ConfigIgnoreAutoRoutes(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES).(bool)
	return
}
func getSettingIp4ConfigIgnoreAutoDns(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS).(bool)
	return
}
func getSettingIp4ConfigDhcpClientId(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID).(string)
	return
}
func getSettingIp4ConfigDhcpSendHostname(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME).(bool)
	return
}
func getSettingIp4ConfigDhcpHostname(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME).(string)
	return
}
func getSettingIp4ConfigNeverDefault(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT).(bool)
	return
}
func getSettingIp4ConfigMayFail(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL).(bool)
	return
}

// Setter
func setSettingIp4ConfigMethod(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, value)
}
func setSettingIp4ConfigDns(data _ConnectionData, value []uint32) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, value)
}
func setSettingIp4ConfigDnsSearch(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, value)
}
func setSettingIp4ConfigAddresses(data _ConnectionData, value [][]uint32) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, value)
}
func setSettingIp4ConfigRoutes(data _ConnectionData, value [][]uint32) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, value)
}
func setSettingIp4ConfigIgnoreAutoRoutes(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, value)
}
func setSettingIp4ConfigIgnoreAutoDns(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, value)
}
func setSettingIp4ConfigDhcpClientId(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, value)
}
func setSettingIp4ConfigDhcpSendHostname(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, value)
}
func setSettingIp4ConfigDhcpHostname(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, value)
}
func setSettingIp4ConfigNeverDefault(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, value)
}
func setSettingIp4ConfigMayFail(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL, value)
}

// JSON Getter
func getSettingIp4ConfigMethodJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_METHOD))
	return
}
func getSettingIp4ConfigDnsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS))
	return
}
func getSettingIp4ConfigDnsSearchJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS_SEARCH))
	return
}
func getSettingIp4ConfigAddressesJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ADDRESSES))
	return
}
func getSettingIp4ConfigRoutesJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ROUTES))
	return
}
func getSettingIp4ConfigIgnoreAutoRoutesJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES))
	return
}
func getSettingIp4ConfigIgnoreAutoDnsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS))
	return
}
func getSettingIp4ConfigDhcpClientIdJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID))
	return
}
func getSettingIp4ConfigDhcpSendHostnameJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME))
	return
}
func getSettingIp4ConfigDhcpHostnameJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME))
	return
}
func getSettingIp4ConfigNeverDefaultJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_NEVER_DEFAULT))
	return
}
func getSettingIp4ConfigMayFailJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_MAY_FAIL))
	return
}

// JSON Setter
func setSettingIp4ConfigMethodJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_METHOD))
}
func setSettingIp4ConfigDnsJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS))
}
func setSettingIp4ConfigDnsSearchJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS_SEARCH))
}
func setSettingIp4ConfigAddressesJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ADDRESSES))
}
func setSettingIp4ConfigRoutesJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ROUTES))
}
func setSettingIp4ConfigIgnoreAutoRoutesJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES))
}
func setSettingIp4ConfigIgnoreAutoDnsJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS))
}
func setSettingIp4ConfigDhcpClientIdJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID))
}
func setSettingIp4ConfigDhcpSendHostnameJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME))
}
func setSettingIp4ConfigDhcpHostnameJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME))
}
func setSettingIp4ConfigNeverDefaultJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_NEVER_DEFAULT))
}
func setSettingIp4ConfigMayFailJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_MAY_FAIL))
}

// Remover
func removeSettingIp4ConfigMethod(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD)
}
func removeSettingIp4ConfigDns(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS)
}
func removeSettingIp4ConfigDnsSearch(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH)
}
func removeSettingIp4ConfigAddresses(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES)
}
func removeSettingIp4ConfigRoutes(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES)
}
func removeSettingIp4ConfigIgnoreAutoRoutes(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES)
}
func removeSettingIp4ConfigIgnoreAutoDns(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS)
}
func removeSettingIp4ConfigDhcpClientId(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID)
}
func removeSettingIp4ConfigDhcpSendHostname(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME)
}
func removeSettingIp4ConfigDhcpHostname(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME)
}
func removeSettingIp4ConfigNeverDefault(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT)
}
func removeSettingIp4ConfigMayFail(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL)
}
