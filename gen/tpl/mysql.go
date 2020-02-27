package tpl

var MysqlTpl = `
package {{.Mysql.Pkg}}

import (
	"fmt"
	"{{.Config.Import}}"
	"github.com/jinzhu/gorm"
	"time"
)
{{if .Mysql.Sources}}
var ( {{range .Mysql.Sources}}
	// {{.Annotation}}
	{{upper .Name}} *gorm.DB{{end}}
){{end}}

func Init() (err error) { {{range .Mysql.Sources}}
	{{upper .Name}}, err = initClient("{{.Name}}")
	if err != nil {
		return
	}{{end}}
	return
}

func initClient(key string) (db *gorm.DB, err error) {
	dsn := {{.Config.Pkg}}.Viper.GetString("mysql." + key + ".dsn")
	driver := {{.Config.Pkg}}.Viper.GetString("mysql." + key + ".driver")
	maxIdleConn := {{.Config.Pkg}}.Viper.GetInt("mysql." + key + ".max_idle_conn")
	maxOpenConn := {{.Config.Pkg}}.Viper.GetInt("mysql." + key + ".max_open_conn")
	maxLifeTime := {{.Config.Pkg}}.Viper.GetInt("mysql." + key + ".max_life_time")
	db, err = gorm.Open(driver, dsn)
	if err != nil {
		fmt.Println(fmt.Sprintf("mysql init error, is %s", err.Error()))
		return nil, err
	}
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(maxLifeTime))
	db.DB().SetMaxOpenConns(maxOpenConn)
	db.DB().SetMaxIdleConns(maxIdleConn)
	if {{.Config.Pkg}}.Env == "beta" || {{.Config.Pkg}}.Env == "dev" {
		db.LogMode(true)
	}
	return db, nil
}
`
