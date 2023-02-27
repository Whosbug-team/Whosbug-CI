package analyze

import "errors"

var (
	// ErrAlreadyExists release已存在，且latest-commit未更新
	//	@update 2023-02-28 02:53:49
	ErrAlreadyExists = errors.New("The Project and Release already exist")

	// ErrUnsupportedLanguage 不支持静态解析的语言
	//	@update 2023-02-28 02:53:45
	ErrUnsupportedLanguage = errors.New("Unsupported Language for ast parse")
)
