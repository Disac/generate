package tpl

var ConfigFileTomlTpl = `#配置文件
{{if .Provider.Mysql.Namespace}}{{$mysql_namespace := .Provider.Mysql.Namespace}}[{{$mysql_namespace}}]{{range .Provider.Mysql.Sources}}
	[{{$mysql_namespace}}.{{.Name}}]
        driver = "{{.Driver}}"
        dsn = "{{.Dsn}}"
        max_idle_conn = {{.MaxIdleConn}}
        max_open_conn = {{.MaxOpenConn}}
        max_life_time = {{.MaxLifeTime}}{{end}}{{end}}
{{if .Provider.Redis.Namespace}}{{$redis_namespace := .Provider.Redis.Namespace}}[{{$redis_namespace}}]{{range .Provider.Redis.Sources}}
    [{{$redis_namespace}}.{{.Name}}]
        addr = "{{.Addr}}"
        pwd = "{{.Pwd}}"
        db = {{.Db}}
        pool_size = {{.PoolSize}}{{end}}{{end}}
{{if .Provider.Rabbitmq.Namespace}}{{$rabbitmq_namespace := .Provider.Rabbitmq.Namespace}}[{{$rabbitmq_namespace}}]{{if .Provider.Rabbitmq.Sources.Publishers}}
	[{{$rabbitmq_namespace}}.publisher] {{range .Provider.Rabbitmq.Sources.Publishers}}
		# {{.Annotation}}
    	[{{$rabbitmq_namespace}}.publisher.{{.Name}}]
			url = "{{.URL}}"
			queue = "{{.Queue}}"
			exchange = "{{.Exchange}}"{{end}}{{end}}{{if .Provider.Rabbitmq.Sources.Consumers}}
	[{{$rabbitmq_namespace}}.consumer] {{range .Provider.Rabbitmq.Sources.Consumers}}
		# {{.Annotation}}
    	[{{$rabbitmq_namespace}}.consumer.{{.Name}}]
			url = "{{.URL}}"
			queue = "{{.Queue}}"
			exchange = "{{.Exchange}}"{{end}}{{end}}{{end}}
{{if .Provider.Kafka.Namespace}}{{$kafka_namespace := .Provider.Kafka.Namespace}}[{{$kafka_namespace}}]{{if .Provider.Kafka.Sources.Producers}}
	[{{$kafka_namespace}}.producer] {{range .Provider.Kafka.Sources.Producers}}
		# {{.Annotation}}
    	[{{$kafka_namespace}}.producer.{{.Name}}]
			hosts = [{{join .Hosts}}]
			topic = "{{.Topic}}"{{end}}{{end}}{{if .Provider.Kafka.Sources.Consumers}}
	[{{$kafka_namespace}}.consumer] {{range .Provider.Kafka.Sources.Consumers}}
		# {{.Annotation}}
    	[{{$kafka_namespace}}.consumer.{{.Name}}]
			hosts = [{{join .Hosts}}]
			topics = [{{join .Topics}}]
			group_id = "{{.GroupID}}"{{end}}{{end}}{{end}}
`

var ConfigFileJsonTpl = ``
var ConfigFileYamlTpl = ``
