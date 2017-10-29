# Changelog
## [v0.6.0]
### Added 
#### WNE-601-lbs-maclists
 - Add MAC-filters for LBS
#### WNE-604-clients-unique
 - Add fields for connection to WirelessClient
#### WNE-584-maxclients
 - Add ClientSessionInfo struct
#### WNE-583-roaming-settings
 - PMK caching and 802.11r settings
#### WNE-591-back-scan-table
 - Add BackScanModelRaw for handle data from CPE
#### capability-band
 - Frequency band added to capabilities
#### WNE-607-l2tp-tunnels
 - Add Models for L2TP tunnels and corresponding CPE settings
#### WNE-568-session-info
 - Add model for collect cpe_session_info
 - Add model for collect client_session_info
#### WNE-603-JSON-RPC-primitives
 - Add JSON RPC primitives
 - Add Tunnel Manager support
 - Tunnel Manager broadcast meta
 - Add RPC structures for Tunnel Manager

### Changed
#### WNE-630-Add-new-fields-to-client-model
 - Add channel parameters (rssi, noise, inactive, mode) to Client and Client Info 
 - Add channel paramters and client hw modes to ClientStat
#### Fix-Backgound-Scanning-JSON
 - Model fixes for BackScanModelRaw

## [v0.5.0] - 09-10-2017
### Added 
#### WNE-510-CPE-Stats
 - Add model for collect "cpe_stat_info"
 - Add model for collect "wlan_stat_info"
#### WNE-511-new-config-rsp
 - Add struct for smart CPE/Config response with statuses
#### WNE-531-access-controller
 - Add Radius state isLocal - for Beeline intergation
 - Add Event for changing WLAN status at radius changing
 - Access control module
#### WNE-519-wpa1
 - Add WPA1 security inda WLAN
#### WNE-521-reg-domain
 - Add radio country to model
#### WNE-528-cpe-addresses
 - IP/MAC adresses to CPE model
#### WNE-527-radio-to-acct
 - Radio ID/ Freq to client statistics
#### WNE-484-netpacket-structs
 - Add models for DHCP packets send

### Changed
#### WNE-307-cpe-config-and-state
 - Separate config from state in CPE model
#### WNE-564-stat-daemon-growth
 - Changed model for WirelessClient
#### Other
 - Refuctoring of models
 - Bugofixes and improvements of BSON marshallers

## [v0.4.0] - 11-09-2017
### Added
#### WNE-470-request-topics
- В топик запроса добавлена обязательная секция тэга (может быть пустой)
#### WNE-493-cpe-status-updating
- Добавлено состояние CPE "Обновление"
#### WNE-453-heatmap-cpe-module
- В модель LBS-данных добавлено поле BestCPE
- Добавлено значение enum'а для модуля CPE_STAT
#### cpe_config_success
- Добавлено значение enum'а для ответа CPE об успехе конфигурирования

### Changed
#### WNE-469-enum-refactor
- Enum'ы реализованы на основе строковых констант


## [v0.3.0] - 14-08-2017
### Added
#### WNE-298
 - L2 isolation option in WLAN model
#### WNE-362
 - List of allowed channels in WifiData
#### WNE-368
 - Name in stat rules
#### WNE-369
 - Wireless client model
#### WNE-377
 - Common system event model
 - CPE_CONNECTED
 - CPE_DISCONNECTED
 - SERVICE_CONNECTED
 - MONITOR_RULE_VIOLATION
 - CPE_CONFIGURATION_ERROR
 - CLIENT_CONNECTED
 - CLIENT_DISCONNECTED
 - SERVICE_SETTINGS_CHANGE
#### WNE-405
 - MQTTLog module and multilevel wildcard
#### WNE-409
 - Service start function and version object
#### WNE-413
 - Statistics settings in CPE
#### WNE-415
 - RequestTopicWithTag
#### WNE-420
 - SystemEventAny
#### WNE-421
 - Log settings in CPE model
#### WNE-422
 - Accounting interval in WLAN model
#### Other
 - CPE id and WLAN id in accounting stat
 - MonitorRuleViolationData.Stat_id

### Changed
#### WNE-367
 - Change radius model to be always accounting and authentication
#### WNE-377
 - BSON fixes for events
#### WNE-416
 - Allow multiple subscription in wrappers
#### WNE-417
 - Use 'timestamp' json key
#### WNE-418
 - Topic parsers return non-pointer Topic and error
#### WNE-426
 - Check if associated data is null in JSON
#### Other
 - `cpeid` -> `cpe_id` in JSON
 - More informative enum serialization errors

### Removed
#### WNE-377
 - Stat rule violation event
 - CPE Connected
 - CPE Disconnected
 - MonitorRuleViolationData.Rule_name

## [v0.0.2] - 17-07-2017
### Added
#### WNE-265 / WNE-282
* Capabilities models
#### WNE-276
* LBS models
#### Other
* Add Model::Dummy
* Add extra MQTT utilities

### Changed
#### WNE-297
* Update CPE model for LBS configuration
#### Other
* MQTTDocummentMessage.D accepts any type as payload now
* Change timestamp field name in utils
