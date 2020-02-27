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
{{if .Rabbitmq.Namespace}}{{$rabbitmq_namespace := .Rabbitmq.Namespace}}[{{$rabbitmq_namespace}}]{{if .Rabbitmq.Sources.Publishers}}
	[{{$rabbitmq_namespace}}.publisher] {{range .Rabbitmq.Sources.Publishers}}
		# {{.Annotation}}
    	[{{$rabbitmq_namespace}}.publisher.{{.Name}}]
			url = "{{.URL}}"
			queue = "{{.Queue}}"
			exchange = "{{.Exchange}}"{{end}}{{end}}{{if .Rabbitmq.Sources.Consumers}}
	[{{$rabbitmq_namespace}}.consumer] {{range .Rabbitmq.Sources.Consumers}}
		# {{.Annotation}}
    	[{{$rabbitmq_namespace}}.consumer.{{.Name}}]
			url = "{{.URL}}"
			queue = "{{.Queue}}"
			exchange = "{{.Exchange}}"{{end}}{{end}}{{end}}
`

var ConfigFileJsonTpl = ``
var ConfigFileYamlTpl = ``
