package gen

import (
	"fmt"
	"git-admin.inyuapp.com/feature/generate/gen/tpl"
	"log"
	"os"
	"text/template"
)

const TplNameMysqlCode = "mysql_code"

func (g *Generator) MysqlCode() {
	path := g.ProviderDir + g.Mysql.Pkg
	err := g.mkdirAll(path)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	t, err := template.New(TplNameMysqlCode).Funcs(fns).Parse(tpl.MysqlCodeTpl)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f, _ := os.Create(fmt.Sprintf(FileFormat, path, g.Mysql.Pkg, "go"))
	err = t.Execute(f, g)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f.Close()
	g.cmd(path)
	return
}
