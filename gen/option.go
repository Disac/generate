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