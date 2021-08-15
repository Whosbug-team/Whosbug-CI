package antlrpack

type JavaTreeShapeListener struct {
	Infos AnalysisInfoType
}

type KotlinTreeShapeListener struct {
	MethodInfo MethodInfoType
	Infos      AnalysisInfoType
}

type CppTreeShapeListener struct {
	Type        string
	Declaration []MemberType
	Infos       AnalysisInfoType
}

type GoTreeShapeListener struct {
	Declaration []MemberType
	Infos       AnalysisInfoType
}

type MemberType struct {
	Name string
	Type string
}
type JSTreeShapeListener struct {
	ObjectInfo ObjectInfoType
	ClassInfo  classInfoType
	Infos      AnalysisInfoType
}
type CallMethodType struct {
	StartLine int
	Id        string
}

type AnalysisInfoType struct {
	CallMethods []CallMethodType
	AstInfoList astInfoType
}

type astInfoType struct {
	Classes []classInfoType
	Methods []MethodInfoType
	Objects []ObjectInfoType
}

type classInfoType struct {
	StartLine    int
	EndLine      int
	ClassName    string
	MasterObject masterObjectInfoType
	Extends      string
}

type MethodInfoType struct {
	StartLine    int
	EndLine      int
	MethodName   string
	MasterObject masterObjectInfoType
	CallMethods  []string
	Params       []string
}

type masterObjectInfoType struct {
	ObjectName string
	StartLine  int
}

type ObjectInfoType struct {
	StartLine   int
	EndLine     int
	ObjectName  string
	ObjFuncName []string
}
