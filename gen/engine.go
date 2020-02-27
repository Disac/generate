package gen

import (
	"fmt"
	"git-admin.inyuapp.com/feature/generate/gen/tpl"
)

const (
	main      = "main"
	publisher = "publisher"
	producer  = "producer"
	consumer  = "consumer"
)

const (
	TplNameMain                  = "main_code"
	TplNameConfigFile            = "config_file"
	TplNameConfigParseCode       = "config_parse_code"
	TplNameMysqlCode             = "mysql_code"
	TplNameRedisCode             = "redis_code"
	TplNameRabbitmqCode          = "rabbitmq_code"
	TplNameRabbitmqPublisherCode = "rabbitmq_publisher_code"
	TplNameRabbitmqConsumerCode  = "rabbitmq_consumer_code"
	TplNameKafkaProducerCode     = "kafka_producer_code"
	TplNameKafkaConsumerCode     = "kafka_consumer_code"
)

func (g *Generator) ConfigFile() {
	path := g.root + g.Config.Path
	text := ""
	switch g.Config.Type {
	case "json":
		text = tpl.ConfigFileJsonTpl
	case "yaml":
		text = tpl.ConfigFileYamlTpl
	default:
		text = tpl.ConfigFileTomlTpl
	}
	fileName := fmt.Sprintf(FileFormat, path, "dev", g.Config.Type)
	g.generate(path, TplNameConfigFile, text, fileName, g)
	return
}

func (g *Generator) ConfigParseCode() {
	path := g.root + g.Config.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Config.Pkg, "go")
	g.generate(path, TplNameConfigParseCode, tpl.ConfigParseTpl, fileName, g.Config)
	return
}

func (g *Generator) MysqlCode() {
	g.Provider.Mysql.Import = fmt.Sprintf(ImportFormat, g.Provider.Import, g.Provider.Mysql.Pkg)
	path := g.Provider.Dir + g.Provider.Mysql.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Provider.Mysql.Pkg, "go")
	g.generate(path, TplNameMysqlCode, tpl.MysqlTpl, fileName, g)
	return
}

func (g *Generator) RedisCode() {
	g.Provider.Redis.Import = fmt.Sprintf(ImportFormat, g.Provider.Import, g.Provider.Redis.Pkg)
	path := g.Provider.Dir + g.Provider.Redis.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Provider.Redis.Pkg, "go")
	g.generate(path, TplNameRedisCode, tpl.RedisTpl, fileName, g)
	return
}

func (g *Generator) InitRabbitmqCode() {
	g.Provider.Rabbitmq.Import = fmt.Sprintf(ImportFormat, g.Provider.Import, g.Provider.Rabbitmq.Pkg)
	path := g.Provider.Dir + g.Provider.Rabbitmq.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Provider.Rabbitmq.Pkg, "go")
	g.generate(path, TplNameRabbitmqCode, tpl.RabbitmqTpl, fileName, g.Provider.Rabbitmq)
}

func (g *Generator) RabbitmqPublisherCode() {
	path := g.Provider.Dir + g.Provider.Rabbitmq.Pkg
	fileName := fmt.Sprintf(FileFormat, path, publisher, "go")
	g.generate(path, TplNameRabbitmqPublisherCode, tpl.RabbitmqPublisherTpl, fileName, g)
	return
}

func (g *Generator) RabbitmqConsumerCode() {
	path := g.Provider.Dir + g.Provider.Rabbitmq.Pkg
	fileName := fmt.Sprintf(FileFormat, path, consumer, "go")
	g.generate(path, TplNameRabbitmqConsumerCode, tpl.RabbitmqConsumerTpl, fileName, g)
	return
}

func (g *Generator) InitKafkaCode() {
	g.Provider.Kafka.Import = fmt.Sprintf(ImportFormat, g.Provider.Import, g.Provider.Kafka.Pkg)
}

func (g *Generator) KafkaProducerCode() {
	path := g.Provider.Dir + g.Provider.Kafka.Pkg
	fileName := fmt.Sprintf(FileFormat, path, producer, "go")
	g.generate(path, TplNameKafkaProducerCode, tpl.KafkaProducerTpl, fileName, g)
	return
}

func (g *Generator) KafkaConsumerCode() {
	path := g.Provider.Dir + g.Provider.Kafka.Pkg
	fileName := fmt.Sprintf(FileFormat, path, consumer, "go")
	g.generate(path, TplNameKafkaConsumerCode, tpl.KafkaConsumerTpl, fileName, g)
	return
}

func (g *Generator) MainCode() {
	path := g.root
	fileName := fmt.Sprintf(FileFormat, g.root, main, "go")
	g.generate(path, TplNameMain, tpl.MainTpl, fileName, g)
	return
}
