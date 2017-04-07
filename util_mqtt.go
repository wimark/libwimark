package libwimark

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func MQTTConnectSync(addr string) (mqtt.Client, error) {
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(addr)

	var client = mqtt.NewClient(opts)
	var token = client.Connect()
	token.Wait()

	return client, token.Error()
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
	Payload() []byte
	Retained() bool
}

type MQTTDocumentMessage struct {
	T Topic
	D map[string]interface{}
	R bool
}

func (self *MQTTDocumentMessage) Topic() Topic {
	return self.T
}

func (self *MQTTDocumentMessage) Payload() []byte {
	var payload = []byte{}
	if self.D != nil {
		payload, _ = json.Marshal(self.D)
	}
	return payload
}

func (self *MQTTDocumentMessage) Retained() bool {
	return self.R
}

type MQTTRawMessage struct {
	T Topic
	D []byte
	R bool
}

func (self *MQTTRawMessage) Topic() Topic {
	return self.T
}

func (self *MQTTRawMessage) Payload() []byte {
	return self.D
}

func (self *MQTTRawMessage) Retained() bool {
	return self.R
}

func MQTTMakePublishChan(client mqtt.Client, logger *log.Logger) chan<- MQTTMessage {
	var publishChan = make(chan MQTTMessage)
	go func() {
		for msg := range publishChan {
			var topic_str = msg.Topic().TopicPath()
			var payload = msg.Payload()

			logger.Printf("Sending message - Topic: %s, Payload: %s\n", topic_str, payload)
			client.Publish(topic_str, 2, msg.Retained(), payload)
		}
	}()

	return publishChan
}

type MsgCb func(mqtt.Message)

func MQTTMustSubscribeSync(client mqtt.Client, topic Topic, cb MsgCb) {
	var token = client.Subscribe(topic.TopicPath(), 2, func(_ mqtt.Client, msg mqtt.Message) {
		cb(msg)
	})

	token.Wait()
	var err = token.Error()
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
