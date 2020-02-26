package gen

import (
	"encoding/json"
	"fmt"
	"git-admin.inyuapp.com/feature/generate/gen/model"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	ErrFormat = "new gen error: %s"
)

const (
	FileFormat = "%s" + string(os.PathSeparator) + "%s.%s"
)

const (
	DefaultProviderPkg = "provider"
)

var GOPATH string

func init() {
	GOPATH = strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))[0]
}

func NewGenerator(opt ...Option) *Generator {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	g := &Generator{}
	err = json.Unmarshal(b, &g)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	g.init()
	for _, o := range opt {
		o(g)
	}
	return g
}

type Generator struct {
	Project     string `json:"project"`
	Root        string `json:"root"`
	ProviderDir string `json:"provider_dir"`

	Config model.Config `json:"config"`
	Mysql  model.Mysql  `json:"mysql"`
	Redis  model.Redis  `json:"redis"`
}

func (g *Generator) init() {
	g.Root = GOPATH + string(os.PathSeparator) + "src" + string(os.PathSeparator) + g.Project + string(os.PathSeparator)
	err := os.MkdirAll(g.Root, os.ModePerm)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	g.ProviderDir = g.Root + DefaultProviderPkg + string(os.PathSeparator)
	g.Config.Import = g.Project + string(os.PathSeparator) + g.Config.Pkg
}

func (g *Generator) Run() {
	if g.Config.GenParseCode {
		g.ConfigParseCodeFile()
	}
	if g.Config.GenFile {
		g.ConfigFile()
	}
	if g.Mysql.GenCode {
		g.MysqlCode()
	}
	if g.Redis.GenCode {
		g.RedisCode()
	}
}

func (g *Generator) mkdirAll(path ...string) error {
	for _, p := range path {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) cmd(path string) {
	cmd := exec.Command("go", "fmt", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
