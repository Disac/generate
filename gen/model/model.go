package model

type Base struct {
	Namespace string `json:"namespace"`
	Pkg       string `json:"pkg"`
}

type SourceBase struct {
	Name       string `json:"name"`
	Annotation string `json:"annotation"`
}
