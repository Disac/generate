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
	"text/template"
)

const (
	ErrFormat    = "new gen error: %s"
	FileFormat   = "%s" + string(os.PathSeparator) + "%s.%s"
	ImportFormat = "%s" + string(os.PathSeparator) + "%s"
)

const (
	DefaultModelDir    = "models" + string(os.PathSeparator)
	DefaultModuleDir   = "modules" + string(os.PathSeparator)
	DefaultProviderDir = "providers" + string(os.PathSeparator)
)

var GOPATH string

func init() {
	GOPATH = strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))[0]
}

func NewGenerator(path string, opt ...Option) *Generator {
	b, err := ioutil.ReadFile(path)
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
	Project  string         `json:"project"`
	Config   model.Config   `json:"config"`
	Provider model.Provider `json:"provider"`

	root string
	src  string
	dir  model.Dir
}

func (g *Generator) init() {
	g.src = GOPATH + string(os.PathSeparator) + "src" + string(os.PathSeparator)
	g.root = g.src + g.Project + string(os.PathSeparator)
	err := os.MkdirAll(g.root, os.ModePerm)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	g.dir = model.Dir{
		Models:    g.root + DefaultModelDir,
		Modules:   g.root + DefaultModuleDir,
		Providers: g.root + DefaultProviderDir,
	}
	g.Config.Import = g.Project + string(os.PathSeparator) + g.Config.Pkg
	g.Provider.Import = g.Project + string(os.PathSeparator) + g.Provider.Pkg
	g.Provider.Dir = g.root + g.Provider.Pkg + string(os.PathSeparator)
}

func (g *Generator) Run() {
	if g.Config.GenFile {
		g.ConfigFile()
	}
	if g.Config.GenParseCode {
		g.ConfigParseCode()
	}
	if g.Provider.Mysql.GenCode {
		g.MysqlCode()
	}
	if g.Provider.Redis.GenCode {
		g.RedisCode()
	}
	if g.Provider.Rabbitmq.GenPublisherCode {
		g.Provider.Rabbitmq.Once.Do(func() {
			g.InitRabbitmqCode()
		})
		g.RabbitmqPublisherCode()
	}
	if g.Provider.Rabbitmq.GenConsumerCode {
		g.Provider.Rabbitmq.Once.Do(func() {
			g.InitRabbitmqCode()
		})
		g.RabbitmqConsumerCode()
	}
	if g.Provider.Kafka.GenProducerCode {
		g.Provider.Kafka.Once.Do(func() {
			g.InitKafkaCode()
		})
		g.KafkaProducerCode()
	}
	if g.Provider.Kafka.GenConsumerCode {
		g.Provider.Kafka.Once.Do(func() {
			g.InitKafkaCode()
		})
		g.KafkaConsumerCode()
	}
	g.MainCode()
}

func (g *Generator) generate(path, tplName, tpl, fileName string, data interface{}) {
	err := g.mkdirAll(path)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	t, err := template.New(tplName).Funcs(fns).Parse(tpl)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f, _ := os.Create(fileName)
	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	f.Close()
	if strings.HasSuffix(f.Name(), ".go") {
		g.cmd(path)
	}
	return
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
