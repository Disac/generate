package main

import (
	"flag"
	"git-admin.inyuapp.com/feature/generate/gen"
)

var config = flag.String("config", "./config.json", "")

const defaultPath = "./config.json"

func main() {
	flag.Parse()
	path := *config
	if path == "" {
		path = defaultPath
	}
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
