package main

import (
	"flag"
	"git-admin.inyuapp.com/feature/generate/gen"
)

const defaultPath = "./config.json"

var config = flag.String("config", defaultPath, "-config=./config.json")

func main() {
	flag.Parse()
	path := *config
	g := gen.NewGenerator(path,
		//gen.CloseGenerateConfigParseCode(),
		//gen.CloseGenerateConfigFile(),
		//gen.CloseGenerateMysqlCode(),
		//gen.CloseGenerateRedisCode(),
		//gen.CloseGenerateRabbitmqConsumerCode(),
		//gen.CloseGenerateRabbitmqPublisherCode(),
		//gen.CloseGenerateKafkaConsumerCode(),
		//gen.CloseGenerateKafkaProducerCode(),
	)
	g.Run()
}
