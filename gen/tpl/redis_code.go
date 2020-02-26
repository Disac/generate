package tpl

var RedisCodeTpl = `
package {{.Redis.Pkg}}

import (
	"{{.Config.Import}}"
	"github.com/go-redis/redis"
)

{{range .Redis.Sources}}
var {{upper .Name}} *redis.Client{{end}}

func Init() (err error) { {{range .Redis.Sources}}
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
