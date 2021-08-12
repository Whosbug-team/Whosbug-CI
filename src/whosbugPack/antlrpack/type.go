package antlrpack

type JavaTreeShapeListener struct {
	Infos AnalysisInfoType
}

type KotlinTreeShapeListener struct {
	MethodInfo MethodInfoType
	Infos      AnalysisInfoType
}

type GoTreeShapeListener struct{
	InType string
	typeInfo typeInfoType	//这里不直接用structInfoType是考虑到有其他类型的Type，因此typeInfoType存有所有Type的共有属性
	structInfo structInfoType //structInfoType就“继承”typeInfoType
	interfaceInfo interfaceInfoType
	memberInfo memberInfoType //这里主要是同步struct中成员名与成员类型

	Infos AnalysisInfoType
}


type CallMethodType struct {
	StartLine int
	Id        string
}
type AnalysisInfoType struct {
	CallMethods []CallMethodType
	AstInfoList astInfoType
}

