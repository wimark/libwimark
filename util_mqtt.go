package libwimark

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	cache "github.com/patrickmn/go-cache"
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
	return MQTTServiceStartWithId(addr, s, v, "", meta)
}

func MQTTServiceStartWithId(addr string, s Module, v Version, id string, meta interface{}) (mqtt.Client, error) {
	ts := time.Now().Unix()

	// prepare disconnect event for will message
	var opts = mqtt.NewClientOptions().AddBroker(addr)
	eventDisconnetTopic := EventTopic{
		SenderModule: s,
		SenderID:     "",
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
		SenderID:     id,
		Type:         SystemEventTypeServiceConnected,
	}

	level = SystemEventLevelINFO

	eventData := ServiceConnectedData{
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

func MarshalInline(val interface{}) (b []byte, e error) {
	var err error
	var bytes []byte
	bytes, err = json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var v = reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return bytes, nil
	}

	var m = map[string]json.RawMessage{}
	err = json.Unmarshal(bytes, &m)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("inline")
		if len(tag) != 0 {
			msg, e := json.Marshal(v.Field(i).Interface())
			if e != nil {
				return nil, e
			}
			var f = map[string]json.RawMessage{}
			e = json.Unmarshal(msg, &f)
			if e != nil {
				return nil, e
			}
			for n, v := range f {
				m[n] = v
			}
		}
	}

	return json.Marshal(m)
}

func UnmarshalInline(b []byte, val interface{}, tmpl interface{}) error {

	var doc map[string]json.RawMessage
	var err error
	if err = json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	if err = json.Unmarshal(b, val); err != nil {
		return err
	}
	// if 'val' is not a ptr, we'll be kicked out by json.Unmarshal
	var v = reflect.ValueOf(val).Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}

	var ind = -1
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("inline")
		if len(tag) == 0 {
			delete(doc, field.Name)
		} else {
			ind = i
		}
	}
	if ind < 0 {
		return nil
	}
	var bb []byte
	bb, err = json.Marshal(doc)
	var f = v.Field(ind)
	var tv = reflect.ValueOf(tmpl).Elem()
	if f.Kind() != tv.Kind() {
		return errors.New("Inline template is not equal to inline field")
	}
	if err = json.Unmarshal(bb, &tmpl); err != nil {
		return err
	}
	f.Set(tv)

	return nil
}
