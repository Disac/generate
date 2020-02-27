package tpl

var KafkaProducerTpl = `
package {{.Provider.Kafka.Pkg}}

import (
	"encoding/json"
	"{{.Config.Import}}"
	"github.com/Shopify/sarama"
)
{{if .Provider.Kafka.Sources.Producers}}
var ( {{range .Provider.Kafka.Sources.Producers}}
	// {{.Annotation}}
	{{upper .Name}}Producer *Client{{end}}
){{end}}

func InitProducers() (err error) { {{range .Provider.Kafka.Sources.Producers}}
	{{upper .Name}}Producer, err = initClient("{{.Name}}")
	if err != nil {
		return
	}{{end}}
	return
}

func initClient(key string) (client *Client, err error) {
	client = &Client{
		topic: {{.Config.Pkg}}.Viper.GetString("kafka.producer." + key + ".topic"),
	}
	hosts := {{.Config.Pkg}}.Viper.GetStringSlice("kafka.producer." + key + ".hosts")
	client.producer, err = sarama.NewSyncProducer(hosts, producerConfig())
	if err != nil {
		return
	}
	return
}

type Client struct {
	producer sarama.SyncProducer
	topic    string
}

func (client *Client) Send(data interface{}) (err error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: client.topic,
		Value: sarama.StringEncoder(string(dataBytes)),
	}
	_, _, err = client.producer.SendMessage(msg)
	if err != nil {
		return
	}
	return
}

func producerConfig() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	conf.Producer.Retry.Max = 1                    // Retry up to 10 times to produce the message
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V0_11_0_2
	return conf
}
`
var KafkaConsumerTpl = `
package {{.Provider.Kafka.Pkg}}

import (
	"{{.Config.Import}}"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
)
{{if .Provider.Kafka.Sources.Consumers}}
var ( {{range .Provider.Kafka.Sources.Consumers}}
	// {{.Annotation}}
	{{upper .Name}}Consumer *cluster.Consumer{{end}}
){{end}}

func InitConsumers() (err error) { {{range .Provider.Kafka.Sources.Consumers}}
	{{upper .Name}}Consumer, err = initConsumer("{{.Name}}")
	if err != nil {
		return
	}{{end}}
	return
}

func initConsumer(key string) (consumer *cluster.Consumer, err error) {
	hosts:=   {{.Config.Pkg}}.Viper.GetStringSlice("kafka.consumer." + key + ".hosts")
	topics:=  {{.Config.Pkg}}.Viper.GetStringSlice("kafka.consumer." + key + ".topics")
	groupID:= {{.Config.Pkg}}.Viper.GetString("kafka.consumer." + key + ".group_id")
	consumer, err = cluster.NewConsumer(hosts, groupID, topics, consumerConfig())
	if err != nil {
		return
	}
	return
}

func consumerConfig() *cluster.Config {
	conf := cluster.NewConfig()
	conf.Consumer.Return.Errors = true
	conf.Version = sarama.V0_11_0_2
	conf.Group.Return.Notifications = true
	conf.Group.PartitionStrategy = cluster.StrategyRoundRobin
	return conf
}
`
