package main

import "git-admin.inyuapp.com/feature/generate/gen"

func main() {
	g := gen.NewGenerator(
		//gen.CloseGenerateConfigParseCode(),
		//gen.CloseGenerateConfigFile(),
		//gen.CloseGenerateMysqlCode(),
		//gen.CloseGenerateRedisCode(),
	)
	g.Run()
}
