package gen

import (
	"bytes"
	"fmt"
	"git-admin.inyuapp.com/feature/generate/gen/tpl"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const (
	ErrFormat = "new gen error: %s"
)

const (
	FileFormat = "%s" + string(os.PathSeparator) + "%s.go"
)
const (
	defaultConfigPath = "etc/app"
	defaultConfigType = "toml"
)

var GOPATH string

func init() {
	GOPATH = strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))[0]
}
func NewGen(project string, opt ...Option) *Gen {
	if project == "" {
		log.Fatal(fmt.Sprintf(ErrFormat, "projece is nil"))
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(GOPATH)
	buf.WriteByte(os.PathSeparator)
	buf.WriteString("src")
	buf.WriteByte(os.PathSeparator)
	buf.WriteString(strings.Replace(project, "/", string(os.PathSeparator), -1))
	g := &Gen{
		project: project,
		root:    buf.String(),
		pkg:     project[strings.LastIndex(project, "/")+1:],
		config: config{
			code: true,
			path: defaultConfigPath,
			typ:  defaultConfigType,
		},
		code: code{
			config:   true,
			mysql:    true,
			redis:    true,
			rabbitmq: true,
			kafka:    true,
		},
	}
	os.MkdirAll(g.root, os.ModePerm)

	for _, o := range opt {
		o(g)
	}
	return g
}

type Gen struct {
	project string
	root    string
	pkg     string

	config config
	code   code
}

type code struct {
	config   bool
	mysql    bool
	redis    bool
	rabbitmq bool
	kafka    bool
}

type config struct {
	code bool
	path string
	typ  string
}

func (g *Gen) Config() {
	t, err := template.New("default").Parse(tpl.ConfigTpl)
	if err != nil {
		log.Fatal(fmt.Sprintf(ErrFormat, "config tpl err"))
	}
	configPath := g.root + string(os.PathSeparator) + "config"
	os.MkdirAll(configPath, os.ModePerm)
	name := fmt.Sprintf(FileFormat, configPath, "config")
	f, _ := os.Create(name)
	err = t.Execute(f, "")
	if err != nil {
		fmt.Println("a", err)
	}
	f.Close()
	g.cmd(configPath)
	return
}

func (g *Gen) cmd(path string) {
	cmd := exec.Command("go", "fmt", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
