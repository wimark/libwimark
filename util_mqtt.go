package libwimark

import (
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	cache "github.com/patrickmn/go-cache"
	"time"
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

func MQTTServiceStart(addr string, s Module, v Version, meta interface{}) (mqtt.Client, error) {
	ts := time.Now().Unix()

	// prepare disconnect event for will message
	var opts = mqtt.NewClientOptions().AddBroker(addr)
	eventDisconnetTopic := EventTopic{
		SenderModule: s,
		SenderID:     "",
		Type:         SystemEventObjectType{SystemEventServiceDisconnected{}},
	}

	level := SystemEventLevel{SystemEventLevelINFO{}}

	eventObject := SystemEventObject{
		Type: SystemEventObjectType{SystemEventServiceDisconnected{}},
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
		return nil, errors.New(fmt.Sprintf("error: (%s) during service start", err.Error()))
	}
	opts.SetWill(eventDisconnetTopic.TopicPath(), string(b), 2, false)

	opts.SetClientID(s.String())

	client, err := MQTTConnectSyncOpts(opts)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: (%s) while connecting to broker", err.Error()))
	}

	//sending retain status message
	statusTopic := StatusTopic{
		SenderModule: s,
		SenderID:     "",
	}

	statusPayload := ModuleStatus{
		Version: v.Version,
		Commit:  v.Commit,
		Build:   v.Build,
		Service: s,
		Id:      "",
		State:   ServiceState{ServiceStateConnected{}},
		Meta:    meta,
	}

	b, err = json.Marshal(statusPayload)
	if err != nil {
		return client, errors.New(fmt.Sprintf("error: (%s) while marshalling to status message ", err.Error()))
	}
	err = MQTTPublishRetainedSync(client, statusTopic, b)
	if err != nil {
		return client, errors.New(fmt.Sprintf("error: (%s) while publishing retained in service start", err.Error()))
	}

	// sending connect event
	eventConnectTopic := EventTopic{
		SenderModule: s,
		SenderID:     "",
		Type:         SystemEventObjectType{SystemEventServiceConnected{}},
	}

	level = SystemEventLevel{SystemEventLevelINFO{}}

	eventData := ServiceConnectedData{
		Version: v.Version,
		Commit:  v.Commit,
		Build:   v.Build,
	}
	eventObject = SystemEventObject{
		Type: SystemEventObjectType{SystemEventServiceConnected{}},
		Data: eventData,
	}
	eventConnectedPayload := SystemEvent{
		Subject_id: s.String(),
		Timestamp:  ts,
		Level:      level,
	}
	eventConnectedPayload.SystemEventObject = eventObject

	b, err = json.Marshal(eventConnectedPayload)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: (%s) marshalling connected event in service start", err.Error()))
	}

	err = MQTTPublishSync(client, eventConnectTopic, b)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: (%s) publishing connected event in service start", err.Error()))
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

func (self MQTTDocumentMessage) Topic() Topic {
	return self.T
}

func (self MQTTDocumentMessage) Payload() ([]byte, error) {
	if self.D != nil {
		return json.Marshal(self.D)
	} else {
		return nil, nil
	}
}

func (self MQTTDocumentMessage) Retained() bool {
	return self.R
}

type MQTTRawMessage struct {
	T Topic
	D []byte
	R bool
}

func (self MQTTRawMessage) Topic() Topic {
	return self.T
}

func (self MQTTRawMessage) Payload() ([]byte, error) {
	return self.D, nil
}

func (self MQTTRawMessage) Retained() bool {
	return self.R
}

func MQTTPublishMsg(client mqtt.Client, log_cb func(string), msg MQTTMessage) {
	var topic_str = msg.Topic().TopicPath()
	var payload, pErr = msg.Payload()

	if pErr != nil {
		log_cb(fmt.Sprintf("Error marshalling outgoind payload. Topic: %s, Error: %s", topic_str, pErr))
	} else {
		log_cb(fmt.Sprintf("Sending message - Topic: %s, Payload: %s\n", topic_str, payload))
		client.Publish(topic_str, 2, msg.Retained(), payload)
	}
}

func MQTTMakePublishChan(client mqtt.Client, log_cb func(string)) chan<- MQTTMessage {
	var publishChan = make(chan MQTTMessage)
	go func() {
		for msg := range publishChan {
			MQTTPublishMsg(client, log_cb, msg)
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

func MQTTPublishSync(client mqtt.Client, topic Topic, payload interface{}) error {
	var token = client.Publish(topic.TopicPath(), 2, false, payload)
	token.Wait()
	var err = token.Error()
	return err
}

func MQTTPublishRetainedSync(client mqtt.Client, topic Topic, payload interface{}) error {
	var token = client.Publish(topic.TopicPath(), 2, true, payload)
	token.Wait()
	var err = token.Error()
	return err
}

func CacheSetRequest(c *cache.Cache, id string, d time.Duration, v MQTTMessage) {
	c.Set(id, v, d)
}

func CacheGetRequest(c *cache.Cache, id string) MQTTMessage {
	var v, ok = c.Get(id)
	if !ok {
		return nil
	}
	return v.(MQTTMessage)
}
