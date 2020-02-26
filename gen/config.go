package gen

import (
	"fmt"
	"git-admin.inyuapp.com/feature/generate/gen/tpl"
	"log"
	"os"
	"text/template"
)

const TplNameConfigFile = "config_file"
const TplNameConfigParseCode = "config_parse_code"

func (g *Generator) ConfigFile() {
	path := g.Root + g.Config.Path
	err := g.mkdirAll(path)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	text := ""
	switch g.Config.Type {
	case "toml":
		text = tpl.ConfigFileTomlTpl
	}
	t, err := template.New(TplNameConfigFile).Funcs(fns).Parse(text)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f, _ := os.Create(fmt.Sprintf(FileFormat, path, "beta", g.Config.Type))
	err = t.Execute(f, g)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f.Close()
	return
}

func (g *Generator) ConfigParseCodeFile() {
	path := g.Root + g.Config.Pkg
	err := g.mkdirAll(path)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	t, err := template.New(TplNameConfigParseCode).Parse(tpl.ConfigParseCodeTpl)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f, _ := os.Create(fmt.Sprintf(FileFormat, path, g.Config.Pkg, "go"))
	err = t.Execute(f, g.Config)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f.Close()
	g.cmd(path)
	return
}
