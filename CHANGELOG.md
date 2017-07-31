# Changelog
## [v0.0.2] - 17-07-2017
### Added
* WNE-282, WNE-265: Add capabilities models
* WNE-276: Add LBS models
* Add Model::Dummy
* Add extra MQTT utilities

### Changed
* WNE-297: Update CPE model for LBS configuration
* MQTTDocummentMessage.D accepts any type as payload now
* Change timestamp field name in utils

## [v0.3.0] - 11-08-2017
### Added
#### WNE-369
 - Add wirelles client model
#### WNE-377
 - Add common system event model
 - CPE_CONNECTED
 - CPE_DISCONNECTED
 - SERVICE_CONNECTED
 - MONITOR_RULE_VIOLATION
 - CPE_CONFIGURATION_ERROR
 - CLIENT_CONNECTED
 - CLIENT_DISCONNECTED
 - SERVICE_SETTINGS_CHANGE

### Changed
#### WNE-368
 - Add name to stat rules
#### WNE-367
 - Change radius model to be always accounting and authentication 

### Removed
#### WNE-377, 
 - Stat rule violation event
 - CPE Connected
 - CPE Disconnected 
