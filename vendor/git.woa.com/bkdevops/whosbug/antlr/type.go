package antlr

type JavaTreeShapeListener struct {
	AstInfoList astResType
}

type KotlinTreeShapeListener struct {
	AstInfoList astResType
}

type CppTreeShapeListener struct {
	AstInfoList astResType
}

type GoTreeShapeListener struct {
	AstInfoList astResType
}

type JSTreeShapeListener struct {
	AstInfoList astResType
}

type AnalysisInfoType struct {
	AstInfoList astResType
}

type astResType struct {
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
