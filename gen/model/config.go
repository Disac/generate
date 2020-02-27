package model

type Config struct {
	Base
	Path string `json:"path"`
	Type string `json:"type"`

	GenFile      bool `json:"gen_file"`
	GenParseCode bool `json:"gen_parse_code"`
}
