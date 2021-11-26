package libwimark

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MQTTConnectSyncOpts(opts *mqtt.ClientOptions) (mqtt.Client, error) {
	var client = mqtt.NewClient(opts)
	var token = client.Connect()
	token.Wait()

	return client, token.Error()
}

func MQTTMustConnectSyncOpts(opts *mqtt.ClientOptions) mqtt.Client {
	var client, err = MQTTConnectSyncOpts(opts)
	if err != nil {
		panic(err)
	}

	return client
}

func MQTTConnectSync(addr string) (mqtt.Client, error) {
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(addr)
	return MQTTConnectSyncOpts(opts)
}

func MQTTMustConnectSync(addr string) mqtt.Client {
	var client, err = MQTTConnectSync(addr)
	if err != nil {
		panic(err)
	}

	return client
}

func onDisconnect(client mqtt.Client, err error) {
	time.Sleep(time.Second * 1)
	panic("MQTT broker connection lost")
}

func MQTTServiceStart(addr string, s Module, v Version, meta interface{}) (mqtt.Client, error) {
	return MQTTServiceStartWithId(addr, s, v, "", meta)
}

func MQTTServiceStartWithId(addr string, s Module, v Version, id string, meta interface{}) (mqtt.Client, error) {
	ts := time.Now().Unix()

	// prepare disconnect event for will message
	var opts = mqtt.NewClientOptions().AddBroker(addr)
	eventDisconnetTopic := EventTopic{
		SenderModule: s,
		SenderID:     id,
		Type:         SystemEventTypeServiceDisconnected,
	}

	level := SystemEventLevelINFO

	eventObject := SystemEventObject{
		Type: SystemEventTypeServiceDisconnected,
		Data: nil,
	}
	eventDisconnetPayload := SystemEvent{
		Subject_id: s.String(),
		Timestamp:  ts,
		Level:      level,
	}
	eventDisconnetPayload.SystemEventObject = eventObject

	b, err := json.Marshal(eventDisconnetPayload)
	if err != nil {
		return nil, fmt.Errorf("error: (%s) during service start", err.Error())
	}
	opts.SetWill(eventDisconnetTopic.TopicPath(), string(b), 2, false)

	opts.SetClientID(s.String() + id)

	opts.SetConnectionLostHandler(onDisconnect)

	client, err := MQTTConnectSyncOpts(opts)
	if err != nil {
		return nil, fmt.Errorf("error: (%s) while connecting to broker", err.Error())
	}

	//sending retain status message
	statusTopic := StatusTopic{
		SenderModule: s,
		SenderID:     id,
	}

	statusPayload := ModuleStatus{
		Version: v.Version,
		Commit:  v.Commit,
		Build:   v.Build,
		Service: s,
		Id:      id,
		State:   ServiceStateConnected,
		Meta:    meta,
	}

	err = MQTTPublishMsg(client, MQTTDocumentMessage{
		T: statusTopic,
		D: statusPayload,
		R: true,
	})
	if err != nil {
		return client, fmt.Errorf("error: (%s) while publishing retained in service start", err.Error())
	}

	// sending connect event
	eventConnectTopic := EventTopic{
		SenderModule: s,
		SenderID:     id,
		Type:         SystemEventTypeServiceConnected,
	}

	level = SystemEventLevelINFO

	eventData := Version{
		Version: v.Version,
		Commit:  v.Commit,
		Build:   v.Build,
	}
	eventObject = SystemEventObject{
		Type: SystemEventTypeServiceConnected,
		Data: eventData,
	}
	eventConnectedPayload := SystemEvent{
		Subject_id: s.String(),
		Timestamp:  ts,
		Level:      level,
	}
	eventConnectedPayload.SystemEventObject = eventObject

	err = MQTTPublishMsg(client, MQTTDocumentMessage{
		T: eventConnectTopic,
		D: eventConnectedPayload,
		R: false,
	})
	if err != nil {
		return nil, fmt.Errorf("error: (%s) publishing connected event in service start", err.Error())
	}
	return client, nil
}

type MQTTMessage interface {
	Topic() Topic
	Payload() ([]byte, error)
	Retained() bool
}

type MQTTDocumentMessage struct {
	T Topic
	D interface{}
	R bool
}

func (mq MQTTDocumentMessage) Topic() Topic {
	return mq.T
}

func (mq MQTTDocumentMessage) Payload() ([]byte, error) {
	if mq.D != nil {
		return json.Marshal(mq.D)
	} else {
		return nil, nil
	}
}

func (mq MQTTDocumentMessage) Retained() bool {
	return mq.R
}

type MQTTRawMessage struct {
	T Topic
	D []byte
	R bool
}

func (mq MQTTRawMessage) Topic() Topic {
	return mq.T
}

func (mq MQTTRawMessage) Payload() ([]byte, error) {
	return mq.D, nil
}

func (mq MQTTRawMessage) Retained() bool {
	return mq.R
}

func MQTTPublishMsg(client mqtt.Client, msg MQTTMessage) error {
	var topic_str = msg.Topic().TopicPath()
	var payload, pErr = msg.Payload()
	if pErr == nil {
		var token = client.Publish(topic_str, 2, msg.Retained(), payload)
		token.Wait()
		return token.Error()
	}
	return pErr
}

func MQTTMakePublishChan(client mqtt.Client, log_cb func(string)) chan<- MQTTMessage {
	var publishChan = make(chan MQTTMessage)
	go func() {
		for msg := range publishChan {
			var topic_str = msg.Topic().TopicPath()
			var payload, pErr = msg.Payload()
			if pErr != nil {
				log_cb(fmt.Sprintf("Error marshalling outgoind payload. Topic: %s, Error: %s", topic_str, pErr))
			} else {
				log_cb(fmt.Sprintf("Sending message - Topic: %s, Payload: %s\n", topic_str, payload))
				client.Publish(topic_str, 2, msg.Retained(), payload)
			}
		}
	}()

	return publishChan
}

type MsgCb func(mqtt.Message)

func MQTTSubscribeSync(client mqtt.Client, topics []Topic, cb MsgCb) error {
	if topics != nil {
		var topiclist = map[string]byte{}
		for _, t := range topics {
			if t != nil {
				topiclist[t.TopicPath()] = 2
			}
		}

		var token = client.SubscribeMultiple(topiclist, func(_ mqtt.Client, msg mqtt.Message) {
			cb(msg)
		})

		token.Wait()
		var err = token.Error()
		if err != nil {
			return err
		}
	}

	return nil
}

func MQTTMustSubscribeSync(client mqtt.Client, topics []Topic, cb MsgCb) {
	var err = MQTTSubscribeSync(client, topics, cb)
	if err != nil {
		panic(err)
	}
}

func MQTTUnsubscribeSync(client mqtt.Client, topics []Topic) error {
	var err error
	if topics != nil {
		var topiclist = []string{}
		for _, t := range topics {
			if t != nil {
				topiclist = append(topiclist, t.TopicPath())
			}
		}
		var token = client.Unsubscribe(topiclist...)
		token.Wait()
		err = token.Error()
		if err != nil {
			return err
		}
	}

	return err
}

func MQTTMustUnsubscribeSync(client mqtt.Client, topics []Topic) {
	var err = MQTTUnsubscribeSync(client, topics)
	if err != nil {
		panic(err)
	}
}

type LogMsg struct {
	Timestamp time.Time        `json:"timestamp"`
	Level     SystemEventLevel `json:"level"`
	Message   string           `json:"message"`
	Module    Module           `json:"service"`
	ModuleId  UUID             `json:"service_id,omitempty"`
}

type EventWriter struct {
	ch      chan<- MQTTMessage
	makeMsg func(msg LogMsg) MQTTMessage
}

func NewEventWriter(c mqtt.Client, makeMsg func(msg LogMsg) MQTTMessage) EventWriter {
	return EventWriter{ch: MQTTMakePublishChan(c, func(string) {}), makeMsg: makeMsg}
}

func DefaultLoggedEvent(msg LogMsg) MQTTMessage {
	if msg.Level != SystemEventLevelERROR {
		return nil
	}
	return MQTTDocumentMessage{
		T: EventTopic{
			SenderModule: msg.Module,
			SenderID:     string(msg.ModuleId),
			Type:         SystemEventTypeLoggedError,
		},
		D: SystemEvent{
			Timestamp:   msg.Timestamp.Unix(),
			Level:       msg.Level,
			Description: msg.Message,
			SystemEventObject: SystemEventObject{
				Type: SystemEventTypeLoggedError,
				Data: ModelError{
					Module:      msg.Module,
					ModuleId:    msg.ModuleId,
					Type:        WimarkErrorCodeOther,
					Description: msg.Message,
				},
			},
		},
	}
}

func (w EventWriter) Write(p []byte) (n int, err error) {
	var msg LogMsg
	var e = json.Unmarshal(p, &msg)
	if e != nil {
		return 0, e
	}
	var mqttMsg = w.makeMsg(msg)
	if mqttMsg != nil {
		go func() { w.ch <- mqttMsg }()
	}
	return len(p), nil
}
