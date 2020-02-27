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
	ErrFormat  = "new gen error: %s"
	FileFormat = "%s" + string(os.PathSeparator) + "%s.%s"
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
	Project  string         `json:"project"`
	Root     string         `json:"root"`
	Dir      model.Dir      `json:"dir"`
	Config   model.Config   `json:"config"`
	Mysql    model.Mysql    `json:"mysql"`
	Redis    model.Redis    `json:"redis"`
	Rabbitmq model.Rabbitmq `json:"rabbitmq"`
	Kafka    model.Kafka    `json:"kafka"`
}

func (g *Generator) init() {
	g.Root = GOPATH + string(os.PathSeparator) + "src" + string(os.PathSeparator) + g.Project + string(os.PathSeparator)
	err := os.MkdirAll(g.Root, os.ModePerm)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, err))
	}
	g.Dir = model.Dir{
		Models:    g.Root + DefaultModelDir,
		Modules:   g.Root + DefaultModuleDir,
		Providers: g.Root + DefaultProviderDir,
	}
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
	if g.Rabbitmq.GenPublisherCode {
		g.Rabbitmq.Once.Do(func() {
			g.InitRabbitmq()
		})
		g.RabbitmqPublisherCode()
	}
	if g.Rabbitmq.GenConsumerCode {
		g.Rabbitmq.Once.Do(func() {
			g.InitRabbitmq()
		})
		g.RabbitmqConsumerCode()
	}
	if g.Kafka.GenProducerCode {
		g.Kafka.Once.Do(func() {
			g.InitKafka()
		})
		g.KafkaProducerCode()
	}
	if g.Kafka.GenConsumerCode {
		g.Kafka.Once.Do(func() {
			g.InitKafka()
		})
		g.KafkaConsumerCode()
	}
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
