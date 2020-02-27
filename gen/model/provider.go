package model

type Provider struct {
	Base
	Mysql    Mysql    `json:"mysql"`
	Redis    Redis    `json:"redis"`
	Rabbitmq Rabbitmq `json:"rabbitmq"`
	Kafka    Kafka    `json:"kafka"`
}
