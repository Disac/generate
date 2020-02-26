package tpl

var ConfigFileTomlTpl = `#配置文件
{{if .Mysql.Namespace}}{{$mysql_namespace := .Mysql.Namespace}}[{{$mysql_namespace}}]{{range .Mysql.Sources}}
	[{{$mysql_namespace}}.{{.Name}}]
        driver = "{{.Driver}}"
        dsn = "{{.Dsn}}"
        max_idle_conn = {{.MaxIdleConn}}
        max_open_conn = {{.MaxOpenConn}}
        max_life_time = {{.MaxLifeTime}}{{end}}{{end}}
{{if .Redis.Namespace}}{{$redis_namespace := .Redis.Namespace}}[{{$redis_namespace}}]{{range .Redis.Sources}}
    [{{$redis_namespace}}.{{.Name}}]
        addr = "{{.Addr}}"
        pwd = "{{.Pwd}}"
        db = {{.Db}}
        pool_size = {{.PoolSize}}{{end}}{{end}}
`
