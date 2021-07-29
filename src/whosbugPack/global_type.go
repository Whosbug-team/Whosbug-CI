package whosbugPack

// inputJson 存储从input.json读取到的配置信息
type inputJson struct {
	ProjectId       string   `json:"__PROJECT_ID"`
	ReleaseVersion  string   `json:"__RELEASE_VERSION"`
	RepoPath        string   `json:"__PROJECT_URL"`
	BranchName      string   `json:"__BRANCH_NAME"`
	LanguageSupport []string `json:"__LAN_SUPPORT"`
	WebServerHost   string   `json:"__WEB_SRV_HOST"`
}

// commitInfoType 存储每一次commit的信息
type commitInfoType struct {
	startLineNum int
	commitHash   string
	commitEmail  string
	commitName   string
	commitTime   string
}

// changeLineType 存储单个改变行的信息
type changeLineType struct {
	lineNumber int
	changeType string
}

// diffParsedType 解析后的diff信息
type diffParsedType struct {
	diffFileName      string
	diffFilePath      string
	changeLineNumbers []changeLineType
	commitHash        string
	diffContent       map[int]map[string]string
}
