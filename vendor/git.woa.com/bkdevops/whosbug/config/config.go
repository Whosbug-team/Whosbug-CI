package config

import "errors"

// DiffParsedType 解析后的diff信息
type DiffParsedType struct {
	CommitterEmail    string
	CommitTime        string
	CommitAuthor      string
	DiffFileName      string
	ChangeLineNumbers []ChangeLineType
	CommitHash        string
	DiffText          string
	OldLineCount      int
	NewLineCount      int
	TargetLanguage    string
}

// ChangeLineType 存储单个改变行的信息
type ChangeLineType struct {
	// +0.5用于标识移除行位置(避免移出对象范围)
	LineNumber float64
	ChangeType string
}

// ObjectInfoType Ready For New Changes
type ObjectInfoType struct {
	CommitHash string `json:"hash"`
	ID         string `json:"object_id"`
	OldID      string `json:"old_object_id"`
	FilePath   string `json:"path"`
	Parameters string `json:"parameters"`

	StartLine         int `json:"start_line"`
	EndLine           int `json:"end_line"`
	OldLineCount      int `json:"old_line_count"`
	CurrentLineCount  int `json:"current_line_count"`
	RemovedLineCount  int `json:"removed_line_count"`
	AddedNewLineCount int `json:"added_new_line_count"`
}

// Equals 比较两个ObjectInfoType是否相等
//  @receiver s *ObjectInfoType
//  @param b ObjectInfoType
//  @return bool
//  @author: Kevineluo 2022-07-31 07:09:43
func (s *ObjectInfoType) Equals(b ObjectInfoType) bool {
	if s.CommitHash == b.CommitHash && s.ID == b.ID && s.FilePath == b.FilePath && s.StartLine == b.StartLine && s.EndLine == b.EndLine {
		return true
	}
	return false
}

// Config 存储从input.json读取到的配置信息
//  @author: Kevineluo 2022-07-31 07:09:27
type Config struct {
	ProjectURL     string `json:"project_url"`
	ProjectID      string `json:"project_id"`
	ReleaseVersion string `json:"release_version"`
	CryptoKey      string `json:"crypto_key"`

	BranchName    string `json:"branch_name"`
	WebServerHost string `json:"web_server_host"`

	WebServerUserName string `json:"web_server_username"`
	WebServerKey      string `json:"web_server_key"`
}

// 全局变量，存储config信息
var WhosbugConfig Config

// 全局变量，记录起始路径
var WorkPath string

// 全局变量，记录本地最后commit hash
var LocalHashLatest string

// 全局变量，object传输的通道
var ObjectChan chan ObjectInfoType

//// 全局变量，大型object传输的通道
//var ObjectChanLarge chan ObjectInfoType

// SupportLanguagesMap 支持的语言源码文件后缀到语言名的映射
var SupportLanguagesMap = map[string]string{
	// C源码后缀
	".C": "c",
	".c": "c",
	// C++源码后缀
	".cpp": "cpp",
	".cc":  "cpp",
	".cxx": "cpp",
	".hh":  "cpp",
	// Golang源码后缀
	".go": "golang",
	// Java源码后缀
	".java": "java",
	// JavaScript源码后缀
	".js": "javascript",
	// Kotlin源码后缀
	".kt": "kotlin",
}

var LatestCommitHash string

// CommitInfoType 存储每一次commit的信息
type CommitInfoType struct {
	CommitHash     string `json:"hash"`
	CommitterEmail string `json:"email"`
	CommitAuthor   string `json:"author"`
	CommitTime     string `json:"time"`
}

var (
	// AlreadyExistsError release已存在，且lastest-commit未更新
	AlreadyExistsError = errors.New("The Project and Release already exist")
)
