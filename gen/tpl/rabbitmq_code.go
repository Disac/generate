package tpl

var RabbitmqTpl = `
package {{.Rabbitmq.Pkg}}

type Connector struct {
	Url      string
	Queue    string
	Exchange string
}
`

var RabbitmqPublisherCodeTpl = `
package {{.Rabbitmq.Pkg}}

import (
	"{{.Config.Import}}"
	"github.com/assembla/cony"
	"github.com/getsentry/raven-go"
	"github.com/streadway/amqp"
	"net"
	"time"
)
{{if .Rabbitmq.Sources.Publishers}}
var ( {{range .Rabbitmq.Sources.Publishers}}
	// {{.Annotation}}
	{{upper .Name}}Publisher *cony.Publisher{{end}}
){{end}}

func InitPublisher() (err error) { {{range .Rabbitmq.Sources.Publishers}}
	{{upper .Name}}Publisher, err = initPublisher("{{.Name}}")
	if err != nil {
		return
	}{{end}}
	return nil
}

func initPublisher(key string) (publisher *cony.Publisher, err error) {
	connector := Connector{
		Url:      {{.Config.Pkg}}.Viper.GetString("rabbitmq." + key + ".url"),
		Queue:    {{.Config.Pkg}}.Viper.GetString("rabbitmq." + key + ".queue"),
		Exchange: {{.Config.Pkg}}.Viper.GetString("rabbitmq." + key + ".exchange"),
	}
	que := &cony.Queue{
		Name:       connector.Queue, // name
		Durable:    true,            // durable
		AutoDelete: false,           // delete when unused
		Exclusive:  false,           // exclusive
		Args:       nil,
	}
	exchange := cony.Exchange{
		Name:       connector.Exchange,
		Kind:       "topic",
		Durable:    true,
		AutoDelete: false,
		Args:       nil,
	}

	client := cony.NewClient(
		cony.URL(connector.Url),
		cony.Backoff(cony.DefaultBackoff),
		cony.Config(amqp.Config{
			Heartbeat: time.Second * 6,
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
		}),
	)
	client.Declare([]cony.Declaration{
		cony.DeclareExchange(exchange),
		cony.DeclareQueue(que),
		cony.DeclareBinding(cony.Binding{
			Queue:    que,
			Exchange: exchange,
			Key:      "*",
		}),
	})

	publisher = cony.NewPublisher(exchange.Name,
		"*",
		cony.PublishingTemplate(amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "text/json",
		}),
	)
	client.Publish(publisher)
	go func() {
		for client.Loop() {
			select {
			case err := <-client.Errors():
				raven.CaptureError(err, nil)
			}
		}
	}()
	return
}
`

var RabbitmqConsumerCodeTpl = `
package {{.Rabbitmq.Pkg}}

import (
	"{{.Config.Import}}"
	"github.com/assembla/cony"
	"github.com/streadway/amqp"
	"net"
	"time"
)
{{if .Rabbitmq.Sources.Consumers}}
var ( {{range .Rabbitmq.Sources.Consumers}}
	// {{.Annotation}}
	{{upper .Name}}Client   *cony.Client
	{{upper .Name}}Consumer *cony.Consumer{{end}}
){{end}}

func InitConsumer() (err error) { {{range .Rabbitmq.Sources.Consumers}}
	{{upper .Name}}Client, {{upper .Name}}Consumer, err = initConsumer("{{.Name}}")
	if err != nil {
		return
	}{{end}}
	return nil
}

func initConsumer(key string) (client *cony.Client, consumer *cony.Consumer, err error) {
	connector := Connector{
		Url:      {{.Config.Pkg}}.Viper.GetString("rabbitmq." + key + ".url"),
		Queue:    {{.Config.Pkg}}.Viper.GetString("rabbitmq." + key + ".queue"),
		Exchange: {{.Config.Pkg}}.Viper.GetString("rabbitmq." + key + ".exchange"),
	}
	que := &cony.Queue{
		Name:       connector.Queue, // name
		Durable:    true,            // durable
		AutoDelete: false,           // delete when unused
		Exclusive:  false,           // exclusive
		Args:       nil,
	}
	exchange := cony.Exchange{
		Name:    connector.Exchange,
		Kind:    "topic",
		Durable: true,
	}

	client = cony.NewClient(
		cony.URL(connector.Url),
		cony.Backoff(cony.DefaultBackoff),
		cony.Config(amqp.Config{
			Heartbeat: time.Second * 6,
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
		}),
	)
	client.Declare([]cony.Declaration{
		cony.DeclareExchange(exchange),
		cony.DeclareQueue(que),
		cony.DeclareBinding(cony.Binding{
			Queue:    que,
			Exchange: exchange,
			Key:      "*",
		}),
	})
	consumer = cony.NewConsumer(
		que,
		cony.Qos(5000),
	)
	client.Consume(consumer)
	return
}
`
