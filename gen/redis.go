package gen

import (
	"fmt"
	"git-admin.inyuapp.com/feature/generate/gen/tpl"
	"log"
	"os"
	"text/template"
)

const TplNameRedisCode = "redis_code"

func (g *Generator) RedisCode() {
	path := g.Dir.Providers + g.Redis.Pkg
	err := g.mkdirAll(path)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	t, err := template.New(TplNameRedisCode).Funcs(fns).Parse(tpl.RedisCodeTpl)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f, _ := os.Create(fmt.Sprintf(FileFormat, path, g.Redis.Pkg, "go"))
	err = t.Execute(f, g)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f.Close()
	g.cmd(path)
	return
}
