package tpl

var MainTpl = `
package main

import (
	{{if .Config.GenParseCode}}"{{.Config.Import}}"{{end}}
	{{if .Provider.Kafka.Namespace}}"{{.Provider.Kafka.Import}}"{{end}}
	{{if .Provider.Mysql.Namespace}}"{{.Provider.Mysql.Import}}"{{end}}
	{{if .Provider.Redis.Namespace}}"{{.Provider.Redis.Import}}"{{end}}
	{{if .Provider.Rabbitmq.Namespace}}"{{.Provider.Rabbitmq.Import}}"{{end}}
)

func main() { {{if .Config.GenParseCode}}
	err := {{.Config.Pkg}}.Init()
	if err != nil {
		return
	} {{end}}{{if .Provider.Mysql.Namespace}}
	err = {{.Provider.Mysql.Pkg}}.Init()
	if err != nil {
		return
	} {{end}}{{if .Provider.Redis.Namespace}}
	err = {{.Provider.Redis.Pkg}}.Init()
	if err != nil {
		return
	} {{end}}{{if .Provider.Rabbitmq.Namespace}}
	err = {{.Provider.Rabbitmq.Pkg}}.InitPublishers()
	if err != nil {
		return
	} {{end}}{{if .Provider.Kafka.Namespace}}
	err = {{.Provider.Kafka.Pkg}}.InitProducers()
	if err != nil {
		return
	} {{end}}
}

`
