package libwimark

import (
	"encoding/json"
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

func MQTTSubscribeSync(client mqtt.Client, topic Topic, cb MsgCb) error {
	var token = client.Subscribe(topic.TopicPath(), 2, func(_ mqtt.Client, msg mqtt.Message) {
		cb(msg)
	})

	token.Wait()
	var err = token.Error()
	if err != nil {
		return err
	}

	return nil
}

func MQTTMustSubscribeSync(client mqtt.Client, topic Topic, cb MsgCb) {
	var err = MQTTSubscribeSync(client, topic, cb)
	if err != nil {
		panic(err)
	}
}

func MQTTUnsubscribeSync(client mqtt.Client, topic Topic) error {
	var token = client.Unsubscribe(topic.TopicPath())
	token.Wait()
	var err = token.Error()
	if err != nil {
		return err
	}

	return err
}

func MQTTMustUnsubscribeSync(client mqtt.Client, topic Topic) {
	var err = MQTTUnsubscribeSync(client, topic)
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
