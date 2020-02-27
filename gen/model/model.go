package model

type Base struct {
	Namespace string `json:"namespace"`
	Pkg       string `json:"pkg"`
	Import    string `json:"import"`
	Dir       string `json:"dir"`
}

type SourceBase struct {
	Name       string `json:"name"`
	Annotation string `json:"annotation"`
}
