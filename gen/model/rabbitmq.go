package model

type Rabbitmq struct {
	Base
	GenProducerCode bool             `json:"gen_producer_code"`
	GenConsumerCode bool             `json:"gen_consumer_code"`
	Sources         []RabbitmqSource `json:"sources"`
}

type RabbitmqSource struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Topic string `json:"topic"`
	Queue string `json:"queue"`
}
