package gen

type Option func(*Gen)

//ConfigPath 设置config文件路径
func ConfigPath(path, typ string) Option {
	return func(gen *Gen) {
		gen.config.path = path
		gen.config.typ = typ
	}
}

//CloseGenerateConfigCode 关闭生成config代码
func CloseGenerateConfigCode() Option {
	return func(gen *Gen) {
		gen.config.code = false
	}
}

//CloseGenerateMysqlCode 关闭生成mysql代码
func CloseGenerateMysqlCode() Option {
	return func(gen *Gen) {
		gen.code.mysql = false
	}
}

//CloseGenerateRedisCode 关闭生成redis代码
func CloseGenerateRedisCode() Option {
	return func(gen *Gen) {
		gen.code.redis = false
	}
}

//CloseGenerateRabbitmqCode 关闭生成rabbitmq代码
func CloseGenerateRabbitmqCode() Option {
	return func(gen *Gen) {
		gen.code.rabbitmq = false
	}
}

//CloseGenerateKafkaCode 关闭生成kafka代码
func CloseGenerateKafkaCode() Option {
	return func(gen *Gen) {
		gen.code.kafka = false
	}
}
