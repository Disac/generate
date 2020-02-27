package model

import "sync"

type Kafka struct {
	Base
	GenProducerCode bool         `json:"gen_producer_code"`
	GenConsumerCode bool         `json:"gen_consumer_code"`
	Sources         KafkaSources `json:"sources"`

	Once sync.Once
}

type KafkaSources struct {
	Producers []KafkaSource `json:"producers"`
	Consumers []KafkaSource `json:"consumers"`
}

type KafkaSource struct {
	SourceBase
	Hosts   []string `json:"hosts"`
	Topics  []string `json:"topics"`
	Topic   string
	GroupID string `json:"group_id"`
}
