package antlr

type AnalysisInfoType struct {
	AstInfoList AstResType
}

type AstResType struct {
	Classes []ClassInfoType
	Methods []MethodInfoType
}

type ClassInfoType struct {
	StartLine int    `json:"start_line,omitempty"`
	EndLine   int    `json:"end_line,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	Extends   string `json:"extends,omitempty"`
}

type MethodInfoType struct {
	StartLine  int    `json:"start_line,omitempty"`
	EndLine    int    `json:"end_line,omitempty"`
	MethodName string `json:"method_name,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

type LineRangeType struct {
	StartLine int
	EndLine   int
}
