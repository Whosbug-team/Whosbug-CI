package antlr

// AstInfo ast static parse result
//
//	@author kevineluo
//	@update 2023-02-28 12:50:57
type AstInfo struct {
	Classes []Class
	Methods []Method
}

// Class class level object
//
//	@author kevineluo
//	@update 2023-02-28 12:51:37
type Class struct {
	StartLine int    `json:"start_line,omitempty"`
	EndLine   int    `json:"end_line,omitempty"`
	Name      string `json:"name,omitempty"`
	Extends   string `json:"extends,omitempty"`
}

// Method method level object
//
//	@author kevineluo
//	@update 2023-02-28 12:53:10
type Method struct {
	StartLine  int    `json:"start_line,omitempty"`
	EndLine    int    `json:"end_line,omitempty"`
	Name       string `json:"name,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

// LineRange line range
//
//	@author kevineluo
//	@update 2023-02-28 12:53:08
type LineRange struct {
	StartLine int
	EndLine   int
}
