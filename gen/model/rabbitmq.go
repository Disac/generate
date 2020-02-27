package model

import "sync"

type Rabbitmq struct {
	Base
	GenConsumerCode  bool            `json:"gen_consumer_code"`
	GenPublisherCode bool            `json:"gen_publisher_code"`
	Sources          RabbitmqSources `json:"sources"`

	Once sync.Once
}

type RabbitmqSources struct {
	Publishers []RabbitmqSource `json:"publishers"`
	Consumers  []RabbitmqSource `json:"consumers"`
}

type RabbitmqSource struct {
	SourceBase
	URL      string `json:"url"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
}
