package iot

const (
	MessageDownTopic     string = "$oc/devices/{device_id}/sys/messages/down"
	MessageDownTopicName string = "messageDownTopicName"

	MessageUpTopic     string = "$oc/devices/{device_id}/sys/messages/up"
	MessageUpTopicName string = "messageUpTopicName"

	CommandDownTopicName string = "commandDownTopicName"
	CommandDownTopic     string = "$oc/devices/{device_id}/sys/commands/#"

	CommandResponseTopicName string = "commandResponseTopicName"
	CommandResponseTopic     string = "$oc/devices/{device_id}/sys/commands/response/request_id="

	PropertiesUpTopicName string = "propertiesUpTopicName"
	PropertiesUpTopic     string = "$oc/devices/{device_id}/sys/properties/report"

	//平台设置设备属性相关Topic
	PropertiesSetRequestTopicName  string = "propertiesSetRequestTopicName"
	PropertiesSetRequestTopic      string = "$oc/devices/{device_id}/sys/properties/set/#"
	PropertiesSetResponseTopicName string = "propertiesSetResponseTopicName"
	PropertiesSetResponseTopic     string = "$oc/devices/{device_id}/sys/properties/set/response/request_id="

	// 平台查询设备属性
	PropertiesQueryRequestTopicName  string = "propertiesQueryRequestTopicName"
	PropertiesQueryRequestTopic      string = "$oc/devices/{device_id}/sys/properties/get/#"
	PropertiesQueryResponseTopicName string = "propertiesQueryResponseTopicName"
	PropertiesQueryResponseTopic     string = "$oc/devices/{device_id}/sys/properties/get/response/request_id="

	// 设备侧获取平台的设备影子数据
	DeviceShadowQueryRequestTopicName  string = "deviceShadowQueryRequestTopicName"
	DeviceShadowQueryRequestTopic      string = "$oc/devices/{device_id}/sys/shadow/get/request_id="
	DeviceShadowQueryResponseTopicName string = "deviceShadowQueryResponseTopicName"
	DeviceShadowQueryResponseTopic     string = "$oc/devices/{device_id}/sys/shadow/get/response/#"

	// 网关批量上报子设备属性
	GatewayBatchReportSubDeviceTopicName string = "gatewayBatchReportSubDeviceTopicName"
	GatewayBatchReportSubDeviceTopic     string = "$oc/devices/{device_id}/sys/gateway/sub_devices/properties/report"
)
