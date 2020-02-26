package model

type Redis struct {
	Base
	GenCode bool          `json:"gen_code"`
	Sources []RedisSource `json:"sources"`
}

type RedisSource struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Pwd      string `json:"pwd"`
	Db       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}
