package model

type Mysql struct {
	Base
	GenCode bool          `json:"gen_code"`
	Sources []MysqlSource `json:"sources"`
}

type MysqlSource struct {
	SourceBase
	Driver      string `json:"driver"`
	Dsn         string `json:"dsn"`
	MaxIdleConn int    `json:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn"`
	MaxLifeTime int    `json:"max_life_time"`
}
