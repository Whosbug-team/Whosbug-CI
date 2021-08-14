package antlrpack

type JavaTreeShapeListener struct {
	Infos AnalysisInfoType
}

type KotlinTreeShapeListener struct {
	MethodInfo MethodInfoType
	Infos      AnalysisInfoType
}

type CppTreeShapeListener struct{
	Type string
	Declaration []MemberType
	Infos AnalysisInfoType
}

type GoTreeShapeListener struct{
	Infos AnalysisInfoType
}

type MemberType struct {
	Name string
	Value string
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
}

type classInfoType struct {
	StartLine    int
	EndLine      int
	ClassName    string
	MasterObject masterObjectInfoType
}

type MethodInfoType struct {
	StartLine    int
	EndLine      int
	MethodName   string
	MasterObject masterObjectInfoType
	CallMethods  []string
}

type masterObjectInfoType struct {
	ObjectName string
	StartLine  int
}
