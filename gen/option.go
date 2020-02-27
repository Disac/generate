package gen

import "os"

type Option func(*Generator)

//ConfigPath 设置config文件路径，类型，可选包名
func ConfigPath(path, typ string, pkg ...string) Option {
	return func(gen *Generator) {
		gen.Config.Path = path
		gen.Config.Type = typ
		if len(pkg) > 0 {
			gen.Config.Pkg = pkg[0]
			gen.Config.Import = gen.Project + string(os.PathSeparator) + pkg[0]
		}
	}
}

//CloseGenerateConfigParseCode 关闭生成配置解析代码
func CloseGenerateConfigParseCode() Option {
	return func(gen *Generator) {
		gen.Config.GenParseCode = false
	}
}

//CloseGenerateConfigParseCode 关闭生成配置文件
func CloseGenerateConfigFile() Option {
	return func(gen *Generator) {
		gen.Config.GenFile = false
	}
}

//CloseGenerateMysqlCode 关闭生成mysql连接代码
func CloseGenerateMysqlCode() Option {
	return func(gen *Generator) {
		gen.Mysql.GenCode = false
	}
}

//CloseGenerateRedisCode 关闭生成redis连接代码
func CloseGenerateRedisCode() Option {
	return func(gen *Generator) {
		gen.Redis.GenCode = false
	}
}

//CloseGenerateRabbitmqConsumerCode 关闭生成rabbitmq consumer连接代码
func CloseGenerateRabbitmqConsumerCode() Option {
	return func(gen *Generator) {
		gen.Rabbitmq.GenConsumerCode = false
	}
}

//CloseGenerateRabbitmqPublisherCode 关闭生成rabbitmq publisher连接代码
func CloseGenerateRabbitmqPublisherCode() Option {
	return func(gen *Generator) {
		gen.Rabbitmq.GenPublisherCode = false
	}
}

//CloseGenerateKafkaConsumerCode 关闭生成kafka consumer连接代码
func CloseGenerateKafkaConsumerCode() Option {
	return func(gen *Generator) {
		gen.Kafka.GenConsumerCode = false
	}
}

//CloseGenerateKafkaProducerCode 关闭生成kafka producer连接代码
func CloseGenerateKafkaProducerCode() Option {
	return func(gen *Generator) {
		gen.Kafka.GenProducerCode = false
	}
}
