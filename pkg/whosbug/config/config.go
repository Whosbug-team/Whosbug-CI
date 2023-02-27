package config

// Config 存储从input.json读取到的配置信息
//
//	@author: Kevineluo 2022-07-31 07:09:27
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

// WhosbugConfig Config 全局变量，存储config信息
//
//	@update 2023-02-27 11:13:32
var WhosbugConfig Config

// WorkPath string 全局变量，记录起始路径
//
//	@update 2023-02-27 11:13:28
var WorkPath string

// LocalHashLatest string 全局变量，记录本地最后commit hash
//
//	@update 2023-02-27 11:13:36
var LocalHashLatest string

// LatestCommitHash string 已解析记录的最新commit hash
//
//	@update 2023-02-27 11:13:45
var LatestCommitHash string
