package ddDef

type (
	DevLogInOut struct {
		UserName string `json:"username"`
		Ts       int64  `json:"ts"`
		Address  string `json:"addr"`
		ClientID string `json:"clientID"`
		Reason   string `json:"reason"`
	}
	DevPublish struct {
		Ts      int64  `json:"ts"`
		Topic   string `json:"topic"`
		Payload string `json:"payload"`
	}
	InnerPublish struct {
		Topic   string `json:"topic"`
		Payload string `json:"payload"`
	}
)

const (
	// TopicDevPublish dd模块收到设备的发布消息后向内部推送以下topic
	TopicDevPublish = "dd.thing.device.clients.publish"
	// TopicDevLogin dd模块收到设备的登录消息后向内部推送以下topic
	TopicDevLogin = "dd.thing.device.clients.login"
	// TopicDevLogout dd模块收到设备的登出消息后向内部推送以下topic
	TopicDevLogout = "dd.thing.device.clients.logout"
	// TopicInnerPublish dd模块订阅以下topic,收到内部的发布消息后向设备推送
	TopicInnerPublish = "dd.thing.inner.publish"
)
