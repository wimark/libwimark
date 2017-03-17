package libwimark

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MQTTConnectSync(addr string) (mqtt.Client, error) {
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(addr)

	var client = mqtt.NewClient(opts)
	var token = client.Connect()
	token.Wait()

	return client, token.Error()
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
