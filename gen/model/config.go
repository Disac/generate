package model

type Config struct {
	Path   string `json:"path"`
	Type   string `json:"type"`
	Pkg    string `json:"pkg"`
	Import string `json:"import"`

	GenFile      bool `json:"gen_file"`
	GenParseCode bool `json:"gen_parse_code"`
}
