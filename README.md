- [Whosbug 新语法解析/算法记录](#whosbug-%E6%96%B0%E8%AF%AD%E6%B3%95%E8%A7%A3%E6%9E%90%E7%AE%97%E6%B3%95%E8%AE%B0%E5%BD%95)
  - [语法解析调整](#%E8%AF%AD%E6%B3%95%E8%A7%A3%E6%9E%90%E8%B0%83%E6%95%B4)
    - [放弃语法树/方法调用链的解析模式](#%E6%94%BE%E5%BC%83%E8%AF%AD%E6%B3%95%E6%A0%91%E6%96%B9%E6%B3%95%E8%B0%83%E7%94%A8%E9%93%BE%E7%9A%84%E8%A7%A3%E6%9E%90%E6%A8%A1%E5%BC%8F)
    - [改动后，语法解析的目标能力：](#%E6%94%B9%E5%8A%A8%E5%90%8E%E8%AF%AD%E6%B3%95%E8%A7%A3%E6%9E%90%E7%9A%84%E7%9B%AE%E6%A0%87%E8%83%BD%E5%8A%9B)
      - [Java 的变更：](#java-%E7%9A%84%E5%8F%98%E6%9B%B4)
      - [Golang 的变更：](#golang-%E7%9A%84%E5%8F%98%E6%9B%B4)
      - [Kotlin 的变更:](#kotlin-%E7%9A%84%E5%8F%98%E6%9B%B4)
      - [Cpp 的变更：](#cpp-%E7%9A%84%E5%8F%98%E6%9B%B4)
      - [Js 解析调整：](#js-%E8%A7%A3%E6%9E%90%E8%B0%83%E6%95%B4)
    - [数据重整&插件数据流改造](#%E6%95%B0%E6%8D%AE%E9%87%8D%E6%95%B4%E6%8F%92%E4%BB%B6%E6%95%B0%E6%8D%AE%E6%B5%81%E6%94%B9%E9%80%A0)
  - [算法分析调整](#%E7%AE%97%E6%B3%95%E5%88%86%E6%9E%90%E8%B0%83%E6%95%B4)
    - [算法侧可以获得的数据变化：](#%E7%AE%97%E6%B3%95%E4%BE%A7%E5%8F%AF%E4%BB%A5%E8%8E%B7%E5%BE%97%E7%9A%84%E6%95%B0%E6%8D%AE%E5%8F%98%E5%8C%96)
    - [算法的调整：](#%E7%AE%97%E6%B3%95%E7%9A%84%E8%B0%83%E6%95%B4)
    - [算法变更中的收益：](#%E7%AE%97%E6%B3%95%E5%8F%98%E6%9B%B4%E4%B8%AD%E7%9A%84%E6%94%B6%E7%9B%8A)

# Whosbug 新语法解析/算法记录

## 语法解析调整

### 放弃语法树/方法调用链的解析模式

​ 使用 Antlr 解析静态代码语法树存在先天缺陷：**难以解析跨文件调用关系**

​ 解析的粒度调整为：**方法、类的定义行**

​ 方法内的**调用关系调整为仅解析名**

​ Object 的唯一标识方法：path、masterObject、name(包含继承、拓展关系)

​ _TODO:_ 方法重载问题

### 改动后，语法解析的目标能力：

- 解析每个方法、类的声明/定义
- 解析获得每个方法、类的声明包含关系(定义链)

#### Java 的变更：

- Java 解析能力裁剪，只保留获取声明链的方法

改动后的 Java 解析数据结构：

```go
type JavaTreeShapeListener struct {
 	AstInfoList astResType
}

type astResType struct {
	Classes []ClassInfoType
	Methods []MethodInfoType
}

type ClassInfoType struct {
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	ClassName string `json:"class_name"`
	Extends   string `json:"extends"`
}

type MethodInfoType struct {
	StartLine  int    `json:"start_line"`
	EndLine    int    `json:"end_line"`
	MethodName string `json:"method_name"`
	Parameters string `json:"parameters"`
}

```

存储的信息包含：

- 类
  - 类名(包含定义链)
  - 起止行号
  - 继承的父类
- 方法
  - 方法名(包含定义链)
  - 起止行号
  - 参数

手动实现的方法：

- `getParamsOfMethod(ctx *javaparser.MethodDeclarationContext)`
  > 返回方法的参数

```go
func getParamsOfMethod(ctx *javaparser.MethodDeclarationContext) (params string) {
	if ctx.FormalParameters() != nil {
		if ctx.FormalParameters().(*javaparser.FormalParametersContext).FormalParameterList() != nil {
			for index, item := range ctx.FormalParameters().(*javaparser.FormalParametersContext).FormalParameterList().(*javaparser.FormalParameterListContext).AllFormalParameter() {
				params += item.(*javaparser.FormalParameterContext).TypeType().GetText() + " "
				params += item.(*javaparser.FormalParameterContext).VariableDeclaratorId().GetText()
				if index != len(ctx.FormalParameters().(*javaparser.FormalParametersContext).FormalParameterList().(*javaparser.FormalParameterListContext).AllFormalParameter())-1 {
					params += ", "
				}
			}
		}
	}
	return
}
```

- `findJavaDeclarationChain(ctx antlr.ParseTree)`
  > 获取对象的定义链

```go
func findJavaDeclarationChain(ctx antlr.ParseTree) (chainName string) {
	currentContext := ctx.GetParent()
	for {
		if _, ok := currentContext.(*javaparser.ClassDeclarationContext); ok {
			chainName = currentContext.(*javaparser.ClassDeclarationContext).IDENTIFIER().GetText() + "." + chainName
		}
		if _, ok := currentContext.(*javaparser.MethodDeclarationContext); ok {
			chainName = currentContext.(*javaparser.MethodDeclarationContext).IDENTIFIER().GetText() + "." + chainName
		}
		currentContext = currentContext.GetParent()
		if currentContext == nil {
			break
		}
	}
	return
}
```

#### Golang 的变更：

- Golang 解析能力裁剪，只保留获取声明链的能力

改动后的 Golang 解析数据结构:

```diff
type GoTreeShapeListener struct {
	AstInfoList astResType
}

type astResType struct {
	Classes []ClassInfoType
	Methods []MethodInfoType
}

type ClassInfoType struct {
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	ClassName string `json:"class_name"`
	Extends   string `json:"extends"`
}

type MethodInfoType struct {
	StartLine  int    `json:"start_line"`
	EndLine    int    `json:"end_line"`
	MethodName string `json:"method_name"`
	Parameters string `json:"parameters"`
}
```

手动实现的方法：

- `getFunctionAndMethodParams(ctx antlr.ParseTree)`
  > 获得方法或函数的参数

```go
func getFunctionAndMethodParams(ctx antlr.ParseTree) (params string) {
	if _, ok := ctx.(*golang.FunctionDeclContext); ok {
		if ctx.(*golang.FunctionDeclContext).Signature().(*golang.SignatureContext).Parameters() != nil {
			for index, item := range ctx.(*golang.FunctionDeclContext).Signature().(*golang.SignatureContext).Parameters().(*golang.ParametersContext).AllParameterDecl() {
				params += item.(*golang.ParameterDeclContext).IdentifierList().GetText() + " "
				params += item.(*golang.ParameterDeclContext).Type_().GetText()
				if index != len(ctx.(*golang.FunctionDeclContext).Signature().(*golang.SignatureContext).Parameters().(*golang.ParametersContext).AllParameterDecl())-1 {
					params += ", "
				}
			}
		}
	}
	if _, ok := ctx.(*golang.MethodDeclContext); ok {
		if ctx.(*golang.MethodDeclContext).Signature().(*golang.SignatureContext).Parameters() != nil {
			for index, item := range ctx.(*golang.MethodDeclContext).Signature().(*golang.SignatureContext).Parameters().(*golang.ParametersContext).AllParameterDecl() {
				params += item.(*golang.ParameterDeclContext).IdentifierList().GetText() + " "
				params += item.(*golang.ParameterDeclContext).Type_().GetText()
				if index != len(ctx.(*golang.MethodDeclContext).Signature().(*golang.SignatureContext).Parameters().(*golang.ParametersContext).AllParameterDecl())-1 {
					params += ", "
				}
			}
		}
	}
	return
}
```

- `getRecvrTypes(ctx *golang.MethodDeclContext)`
  > 获得方法所属的结构体类型

```go
func getRecvrTypes(ctx *golang.MethodDeclContext) (types []string) {
	temp := ctx.Receiver().(*golang.ReceiverContext).Parameters().(*golang.ParametersContext).AllParameterDecl()
	for _, item := range temp {
		types = append(types, item.(*golang.ParameterDeclContext).Type_().GetText())
	}
	return
}
```

#### Kotlin 的变更:

- Kotlin 解析能力裁剪，只保留获取声明链的能力

改动后的 Kotlin 解析数据结构：

```go
type KotlinTreeShapeListener struct {
	AstInfoList astResType
}

type astResType struct {
	Classes []ClassInfoType
	Methods []MethodInfoType
}

type ClassInfoType struct {
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	ClassName string `json:"class_name"`
	Extends   string `json:"extends"`
}

type MethodInfoType struct {
	StartLine  int    `json:"start_line"`
	EndLine    int    `json:"end_line"`
	MethodName string `json:"method_name"`
	Parameters string `json:"parameters"`
}
```

手动实现的方法：

- `findKotlinClassExtends(ctx *kotlin.ClassDeclarationContext)`
  > 获得类的继承类名

```go
func findKotlinClassExtends(ctx *kotlin.ClassDeclarationContext) (extends string) {
	if ctx.DelegationSpecifiers() != nil {
		tempCtx := ctx.DelegationSpecifiers().(*kotlin.DelegationSpecifiersContext)
		if tempCtx.AllDelegationSpecifier() != nil {
			if tempCtx.DelegationSpecifier(0).(*kotlin.DelegationSpecifierContext).ConstructorInvocation() != nil {
				tempCtx2 := tempCtx.DelegationSpecifier(0).(*kotlin.DelegationSpecifierContext).ConstructorInvocation()
				extends = tempCtx2.(*kotlin.ConstructorInvocationContext).UserType().GetText()
			}
		}
	}
	return
}
```

- `findKotlinDeclarationChain(ctx antlr.ParseTree)`
  > 获得定义链

```go
func findKotlinDeclarationChain(ctx antlr.ParseTree) (chainName string) {
	currentContext := ctx.GetParent()
	for {
		if _, ok := currentContext.(*kotlin.ClassDeclarationContext); ok {
			chainName = currentContext.(*kotlin.ClassDeclarationContext).SimpleIdentifier().GetText() + "." + chainName
		}
		if _, ok := currentContext.(*kotlin.FunctionDeclarationContext); ok {
			chainName = currentContext.(*kotlin.FunctionDeclarationContext).Identifier().GetText() + "." + chainName
		}
		currentContext = currentContext.GetParent()
		if currentContext == nil {
			break
		}
	}
	return
}
```

#### Cpp 的变更：

- Cpp 解析能力裁剪，只保留获取声明链的能力
- 不解析复杂环，只解析简单定义/声明
- 方法名的解析手动实现
  更改后的数据结构：

```go
type CppTreeShapeListener struct {
	AstInfoList astResType
}

type astResType struct {
	Classes []ClassInfoType
	Methods []MethodInfoType
}

type ClassInfoType struct {
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	ClassName string `json:"class_name"`
	Extends   string `json:"extends"`
}

type MethodInfoType struct {
	StartLine  int    `json:"start_line"`
	EndLine    int    `json:"end_line"`
	MethodName string `json:"method_name"`
	Parameters string `json:"parameters"`
}
```

#### Js 解析调整：

Antlr 语法树字段记录：

对于一个 Js 函数：

- 类型：`FunctionDeclarationContext`
- 子字段
  - "function"
  - 函数名 `IdentifierContext`
  - "("
  - 参数 `FormalParameterListContext`
  - ")"
  - 函数体 `FunctionBodyContext`

对于一个 Js 类：

- 类型：`ClassDeclarationContext`
- 子字段
  - "Class"
  - 类名 `IdentifierContext`
  - "("
  - 当含有 extends 字段时，有`ClassTailContext`
  - `SingleExpressionContext`/`ArgumentsExpressionContext`，逐层遍历取`SingleExpression()`可以得到`IdentifierExpressionContext`作为准确的 Extends 名
  - 类内部字段

手动实现的方法：

- `getExtendIdentifier(ctx *javascript.ClassDeclarationContext)(extend string)`

  > 返回类定义中通过 Extends 继承的类名

```go
func getExtendIdentifier(ctx *javascript.ClassDeclarationContext) (extend string) {
	if ctx.ClassTail() != nil {
		tempCtx := ctx.ClassTail().(*javascript.ClassTailContext).SingleExpression()
		for {
			if tempCtx == nil {
				return
			}
			if _, ok := tempCtx.(*javascript.ArgumentsExpressionContext); ok {
				tempCtx = tempCtx.(*javascript.ArgumentsExpressionContext).SingleExpression()
				continue
			}
			if _, ok := tempCtx.(*javascript.IdentifierExpressionContext); ok {
				extend = tempCtx.(*javascript.IdentifierExpressionContext).GetText()
				return
			}
			return
		}
	}
	return
}
```

- `findJsDeclChain(ctx antlr.ParseTree) (chain string)`

  > 返回类/函数定义的定义链关系(所属类、方法等)

```go
func findJsDeclChain(ctx antlr.ParseTree) (chain string) {
	tempCtx := ctx.GetParent()
	for {
		if _, ok := tempCtx.(*javascript.ClassDeclarationContext); ok {
			chain = tempCtx.(*javascript.ClassDeclarationContext).Identifier().GetText() + "." + chain
		}
		if _, ok := tempCtx.(*javascript.FunctionDeclarationContext); ok {
			chain = tempCtx.(*javascript.FunctionDeclarationContext).Identifier().GetText() + "." + chain
		}
		tempCtx = tempCtx.GetParent()
		if tempCtx == nil {
			return
		}
	}
}
```

更改后的数据结构：

```go
type JavascriptTreeShapeListener struct {
	AstInfoList astResType
}

type astResType struct {
	Classes []ClassInfoType
	Methods []MethodInfoType
}

type ClassInfoType struct {
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	ClassName string `json:"class_name"`
	Extends   string `json:"extends"`
}

type MethodInfoType struct {
	StartLine  int    `json:"start_line"`
	EndLine    int    `json:"end_line"`
	MethodName string `json:"method_name"`
	Parameters string `json:"parameters"`
}
```

### 数据重整&插件数据流改造

我们可以看到，经过改造后的几种语言的解析数据结构已经实际统一，我们可以对其进行改造， 共用同一个类型。

改造嵌套类型，函数`antlrAnalysis(diffText string, langMode string) (result astResType)`直接返回`astResType`类型的结构体：

```go
type astResType struct {
	Classes []ClassInfoType
	Methods []MethodInfoType
}
```

`antlrAnalysis`的修改直接影响到`addObjectFromChangeLineNumber()`、`findChangedMethod()`的传入参数类型，需要修改其逻辑/数据流：

```diff
type ObjectInfoType struct {
	CommitHash          string 		`json:"hash"`
	Id                  string 		`json:"object_name"`
	OldId               string 		`json:"old_object_name"`
	FilePath            string 		`json:"path"`
	OldLineCount        int    		`json:"old_line_count"`
	NewLineCount        int    		`json:"new_line_count"`
	ChangedOldLineCount int    		`json:"changed_old_line_count"`
	ChangedNewLineCount int    		`json:"changed_new_line_count"`
+	Type                string 		`json:"type"`
-	Calling				[]string
}
```

## 算法分析调整

> 新的算法模式将重点放置在定义链的解析，对堆栈内容的定义位置做出判断

### 算法侧可以获得的数据变化：

- 减少的数据：
  - Calling 信息不再支持解析
  - MasterObject 不再支持解析
- 新增的数据：
  - 完整的方法/类定义链
  - 起止行号的支持
  - 准确可靠的行数变动情况

### 算法的调整：

- 不再依赖调用链的解析
  > 调用链信息解析困难，且难以提高解析准确性
- 唯一标识对象的模式变更
  > 仅使用定义链+文件目录定位的模式无法准确定位对象，使用定义链条+文件目录+起始行号进行定位
- 置信度计算依赖变更
  > 置信度的计算不再能依赖调用链解析，置信度的维度缩减为依靠时间、行数作为置信度判据

### 算法变更中的收益：

- 不必花费极大的代价实现调用链的准确解析
- **不必等待单次解析的所有 Object 接收完毕再启动算法**
  > 可以在每一个 Object 上传同时进行计算，实时更新置信度
- 插件解析的数据流可以得到简化
  > 插件解析的目标数据大幅度减少，可以缩减数据流中的重复流转信息
- 插件 Antlr 语法解析的效率优化
  > Antlr 解析语法树只需要获得定义信息，可以省去复杂大量的链式分析逻辑
