package gen

import (
	"fmt"
	"git-admin.inyuapp.com/feature/generate/gen/tpl"
)

const (
	publisher = "_publisher"
	consumer  = "_consumer"
)

const (
	TplNameConfigFile            = "config_file"
	TplNameConfigParseCode       = "config_parse_code"
	TplNameMysqlCode             = "mysql_code"
	TplNameRedisCode             = "redis_code"
	TplNameRabbitmqCode          = "rabbitmq_code"
	TplNameRabbitmqPublisherCode = "rabbitmq_publisher_code"
	TplNameRabbitmqConsumerCode  = "rabbitmq_consumer_code"
)

func (g *Generator) ConfigFile() {
	path := g.Root + g.Config.Path
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

func (g *Generator) ConfigParseCodeFile() {
	path := g.Root + g.Config.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Config.Pkg, "go")
	g.generate(path, TplNameConfigParseCode, tpl.ConfigParseCodeTpl, fileName, g.Config)
	return
}

func (g *Generator) MysqlCode() {
	path := g.Dir.Providers + g.Mysql.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Mysql.Pkg, "go")
	g.generate(path, TplNameMysqlCode, tpl.MysqlCodeTpl, fileName, g)
	return
}

func (g *Generator) RedisCode() {
	path := g.Dir.Providers + g.Redis.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Redis.Pkg, "go")
	g.generate(path, TplNameRedisCode, tpl.RedisCodeTpl, fileName, g)
	return
}

func (g *Generator) InitRabbitmq() {
	path := g.Dir.Providers + g.Rabbitmq.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Rabbitmq.Pkg, "go")
	g.generate(path, TplNameRabbitmqCode, tpl.RabbitmqTpl, fileName, g)
}

func (g *Generator) RabbitmqPublisherCode() {
	path := g.Dir.Providers + g.Rabbitmq.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Rabbitmq.Pkg+publisher, "go")
	g.generate(path, TplNameRabbitmqPublisherCode, tpl.RabbitmqPublisherCodeTpl, fileName, g)
	return
}

func (g *Generator) RabbitmqConsumerCode() {
	path := g.Dir.Providers + g.Rabbitmq.Pkg
	fileName := fmt.Sprintf(FileFormat, path, g.Rabbitmq.Pkg+consumer, "go")
	g.generate(path, TplNameRabbitmqConsumerCode, tpl.RabbitmqConsumerCodeTpl, fileName, g)
	return
}
