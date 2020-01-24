# Changelog

## [v1.2.0] - 24-01-2020
### Added
 - staff auth type for MinFin
 - validate field in authentication
 - color schema/popup text for ADs
 - webpage/ad and polls fields updates
 - charts flag and enum for consolidate CPE rep
 - aggrement to page
 - session collections
 - wlan and bss clients cache
 - CPEStatInfo cache variables
 - stats and go mod
 - name/surname to portal request
 - Accounts, Vouchers and tariffs
 - user account / paid plans for paid inet
 - auth/authen types in RedirectRequest
 - traffic shaper config on wired NAT iface
 - ads to watch update for portal ads engine
 - logo/prefix/etc in web
 - MSISDN prefix without DEF codes
 - template/API for CPE disassoc
 - formalize portal client sessions, add config
 - Simple User-account, enums user portal state
 - enums for states of user on portal
 - MSISDN prefix to portal profile
 - preauth list to portal, JSON as basic report
 - report types, enums reformat
 - formalize StatReport obj, add type nums
 - StatReport and StatReportResult for reports
 - support for wired with NAT
 - client username / useragent, new event
 - report formats
 - PortalPageProfile obj, colls for portal
 - CreateAt field to stat
 - portal and coa sessions refac, auth event
 - no masquerade to portal config
 - portal session updates for block support
 - TC per wlan config (models, enum, uci)
 - Nas-Id and fixes for captive portal
 - Nas Port ID to uci to pass in redirect url
 - MACs filter, HTTP/Exec helper funcs
 - structs for radar-export
 - more features for Radar save & analyze
 - common radar API for back as proxy
 - new "new" field in radar for first income
 - visit income for radar analyze
 - station dump data, new Rule types

## [v1.0.7] - 11-01-2019

### Changed
#### lbs-struct-fix
 - Tag in LBS data fixed

## [v1.0.6] - 27-12-2018

### Changed
#### eth-acct-off
 - Write eth_acct section by name
 - Event test fixed
 - Omitempty for eth_acct config section
 - Omitempty support for inline marshaling

## [v1.0.2] - 24-12-2018

### Changed
#### WNE-1147-separate-configure
 - Add vid param to vlan config (filled by configurer)

## [v1.0.1] - 18-12-2018

### Added
#### l3-libtc
 - Add ifaceStats to read users stats
 - Add Ip address as key to libtc
 - libtc: parse interface & users state
 - libtc: user state stuff now is stored by mac
 - libtc: parse U32 filter matches
 - libtc: tc cli driver
 - libtc moved from wbras
#### WNE-1206-acct-mirror
 - Add Accounting mirroring (WLAN, event)
#### cpe-features
 - Cpe.state.firmware.features
#### WNE-1193-mac-auth
 - Use structs from libwimark/libtc
 - Add user classes for QoS from libtc
 - Add QosGroup to redir session
 - Add NasPortId to Acct req
 - Add QosGroup obj
 - Add mac_auth and nas_port_id to WLAN
#### event-desc-root
 - Event description for LOGGED_ERROR type added to root struct
#### beeline-framed-ip-address
 - add framed ip address for beeline configuration purposes
#### WNE-1079-firstboot
 - Default Module value
 - New CPE statuses and last error field
 - Multiple CPEs for reboot + field renames
 - CPE system functions
 - Configurer JSONRPC interface
#### WNE-1077-wmm-support
 - WMM params in UCI config
 - WMM support in models
#### WNE-1075-template-extends
 - Add subnet and MAC prefix to ConfigRule constraints
#### WNE-1081-multiradius
 - Radius section in UCI config
#### WNE-1080-wired-redirect
 - Add UCI config contents to CPE metadata
 - HasCaptiveRedirect in CPE mask
 - Fields for wired redirect
#### WNE-1095-cisco
 - add yaml tags for nesseccary fields
 - Mask for controller
 - CPE mask for controller
 - CPE status meta
 - CPE interface structs moved from configurer
 - Add calls for cisco mediator interface
 - Add model id from CPE to CPE struct
 - Add Controller model for Cisco integration

### Changed
#### WNE-1088-mongo
 - Perf improves in set/get BSON
 - Switch to new mgo
#### WNE-1147-separate-configure
 - Tunnel config moved to cpe state
 - Add Compact objects for frontend using
 - CPE addresses moved to state
 - L2TP config -> state
#### rrm-power
 - New model for Tx power config

## [v0.12.4] - 07-08-2018

### Added
#### WNE-1058-lte-support
 - State of WAN

## [v0.12.3] - 27-07-2018

### Added
#### WNE-1015-hotspot20-settings
 - Generic struct for lang:name
 - Additional HS2.0 profile parameters

## [v0.12.1] - 25-07-2018

### Added
#### RRM-daemon
 - CPE interface for RRM
 - RRM daemon request parameters
 - RRM interface and objects
 - RRMChannel routines
 - Func for list of available Tx powers
#### WNE-1045-req-iface-state
 - New cpe wifi functions added
#### WNE-1040-hetmaps
 - Add OUI vendor list for not-only statdaemon using
#### WNE-1015-hotspot20-settings
 - Support for HS2.0 structs in DB interface
 - Hotspot 2.0 models
#### WNE-1021-monitor-rules-postaction
 - Add post_script and more rule types
#### WNE-997-template-always
 - Add flag for applyng template on every connect

### Changed
#### RRM-daemon
 - Min Tx power in config
 - BSSID in WLAN state
 - Scandata updated: width & htmode
#### WNE-1051-LBS-update
 - Update LBSClientData: freq & float rssi

### Removed
#### mqtt-cleanup
 - Removed unnecessary stuff

## [v0.11.7] - 14-06-2018

### Changed
#### WNE-981-NAT-for-WLAN
 - Option for access to CPE settings via NAT
#### del-client-freq
 - Remove freq from client (dis)connect event data

## [v0.11.3] - 08-06-2018

### Changed
#### template-in-event
 - Add template UUID to CPE connect event
#### enum-verbose
 - Enum errors more verbose

## [v0.11.1] - 05-06-2018

### Added
#### WNE-981-NAT-for-WLAN
 - Additional NAT-related parameters
 - CPE NAT functions
 - NAT flag for WLAN
#### WNE-938-lbs-clients-with-assoc
 - Add fields to save cpe/wlan names in ClientProbe
 - Add manufacturer to client data
 - Fix forgot rssi in struct for lbs
 - Add struct for work with LBS Probe data coll
 - Move LBS objects to stat, add ssid to client obj
#### WNE-882-portal-switch-url
 - Add SwitchURL to portal request object
#### dual-band-support
 - Add frequency (band) to config for frontend needs
#### WNE-794-events-desc
 - Add Description field to SystemEvent
#### WNE-870-redirect-radio-id
 - Add radio id to redirect sessio object
#### cpe-resolved-host
 - Remove unused fields from CPE status
 - Struct for CPE info about broker connection
#### WNE-901-L3ACL
 - Add L3 ACL filter
#### WNE-754-WMM
 - Add DisableWMM flag (so WMM is on by default)

### Changed
#### WNE-983-accounting-refactor
 - Name changing in accounting refactor
 - Change in acct models from cpe
 - CPE-related events refac: one cut short, two sawn off
 - Enums reformatted with nil data in no-value assoc enums
#### WNE-963-cpe-events-refac
 - CPE-related events refac: one cut short, two sawn off
 - Enums reformatted with nil data in no-value assoc enums
#### nil-value-assoc-enum
 - Enums reformatted with nil data in no-value assoc enums
#### WNE-910-error-refac
 - Error struct refac
 - Utility to send logged messages to MQTT
 - Event type for LoggedError
 - JSON inline utilities moved to separate src

### Removed
#### remove-obsolete-caps
 - Remove obsolete capabilities

## [v0.10.17] - 27-04-2018

### Added
#### vlan-option-switch
 - Switch param for VLAN state
#### WNE-864-cpe-package-version
 - Static script versions in CPE state
 - Package versions in CPE firmware state
#### wired-interfaces-to-array-in-db
 - add set get bson to wired interfaces
#### WNE-802-wired-model
 - Switch vlan state

### Changed
#### tunnel-config
 - Reuse Wlan.VLAN for host vlans
 - Added tunnel vlan to wlan model & jsonrpc req
 - HostIP for tunnel manager jsonrpc param
 - HostId tag restored

## [v0.10] - 09-04-2018

### Added
#### marshal-inline
 - Support for multiple inline fields
 - Support for "unique" inline field with constraints
#### WNE-762-portal
 - Models for redirect and redirect API
 - CapriveRedirect / RedirectClientSession / etc
#### WNE-802-wired-model
 - Config, state and capabilities for wired interfaces
 - Interface name for vlan
 - Version for CPEModel
#### WNE-831-accounting-rubic
 - Change API for redirect HTTP requests for RubicPro
#### jsonrpc-funcs
 - JSONRPC support functions (moved from tunnel-manager)
#### WNE-838-broker-reconnect
 - Add onDisconnect function with panic to MQTT util
#### WNE-832-sms-gw
 - Add portal API (status of session, type of authorization)
#### wlan-virtual-iface
 - Field for WLAN virtual interface on CPE
#### cpeagent_jsonrpc
 - New CPEAgent jsonrpc functions
#### wired-client-enum-type
 - Add wired type for WirelessClient
#### IPSec-IP-For-Tun
 - IpSec addr in vpnhost model
 - Add VPN type - openvpn or ipsec
 - add ipsec ip to broadcast
#### jsonrpc-funcs
 - Service start routine with ServiceID support
 - Routines for JSONRPC support

### Changed
#### WNE-731-remote-syslog
 - CPE logging config changed
#### tunnel-config
 - Separate config for tunnels

## [v0.9] - 15-02-2018

### Added
#### WNE-696-fwupdate-event
 - Add events for FW update feature
#### l2chain-template-mask
 - Add l2chain link to mask for config templates

### Changed
#### WNE-744-no-separate-tunnels
 - Tunnel sessions moved from separate coll to CPE model
#### ip-addr-list
 - IP address list for CPE

## [v0.8.0] - 23-01-2018
### Added
#### L2Rules-Template
 - Add L2ACL models
#### WNE-644-jsonrpc-set
 - Support for CPE JSONRPC interface
#### WNE-645-RRM-event
 - Add RRM module
 - Add RRM event
#### WNE-696-firmware-config
 - Add FW sections to CPE state&config
#### WNE-705-fw-update
 - Add FW module
 - Add CPEFirmware struct
#### WNE-709-templatemask-wlan
 - Add mask for template having WLAN

### Changed
#### hotfix0.8.0
 - Tag misprints fixed
 - Defaults for enums

### Removed
#### remove-old-cfg
 - Removed old CPE config/state

## [v0.7.0] - 01-12-2017
### Added
#### WNE-640-templates
 - Add config templates
 - Add config rules
#### WNE-640-split-caps
 - Add cpe model
#### inline-json
 - Add functions for inline support in json marshal/unmarshal
#### WNE-585-mcs-management
 - Add fields for MCS config

### Changed
#### WNE-640-templates
 - Separate capabilities from CPE model
#### new-enums
 - New enum-generator is used
 - Generated enum code updated
 - Removed obsolete enums
#### no-omitemply
 - 'omitempty' tags removed from models

## [v0.6.0] - 03-11-2017
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
#### WNE-614-client-history
 - Extend WirelessClient model to collect RSSI/mode/etc
#### WNE-628-scanning-config
 - Add background scanning config to CPE model

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
