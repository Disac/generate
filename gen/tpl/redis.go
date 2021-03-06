package tpl

var RedisTpl = `
package {{.Provider.Redis.Pkg}}

import (
	"{{.Config.Import}}"
	"github.com/go-redis/redis"
)
{{if .Provider.Redis.Sources}}
var ({{range .Provider.Redis.Sources}}
	// {{.Annotation}}
	{{upper .Name}} *redis.Client{{end}}
){{end}}

func Init() (err error) { {{range .Provider.Redis.Sources}}
	{{upper .Name}}, err = initClient("{{.Name}}")
	if err != nil {
		return
	}{{end}}
	return
}

func initClient(key string) (client *redis.Client, err error) {
	addr := {{.Config.Pkg}}.Viper.GetString("redis." + key + ".addr")
	pwd := {{.Config.Pkg}}.Viper.GetString("redis." + key + ".pwd")
	db := {{.Config.Pkg}}.Viper.GetInt("redis." + key + ".db")
	poolSize := {{.Config.Pkg}}.Viper.GetInt("redis." + key + ".pool_size")

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
		PoolSize: poolSize,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return
	}
	return
}
`
