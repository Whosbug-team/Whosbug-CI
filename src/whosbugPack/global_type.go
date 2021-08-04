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
	commitHash     string
	committerEmail string
	committerName  string
	commitTime     string
}

// changeLineType 存储单个改变行的信息
type changeLineType struct {
	lineNumber int
	changeType string
}

// diffParsedType 解析后的diff信息
type diffParsedType struct {
	committerEmail    string
	commitTime        string
	committerName     string
	diffFileName      string
	changeLineNumbers []changeLineType
	commitHash        string
	diffText          string
	diffContent       map[int]ObjectInfoType
}

// ObjectInfoType 生成的object信息
type ObjectInfoType map[string]string

type postData struct {
	Objects []struct {
		CommitTime string `json:"commit_time"`
		FilePath   string `json:"file_path"`
		Hash       string `json:"hash"`
		Name       string `json:"name"`
		OldName    string `json:"old_name"`
		Owner      string `json:"owner"`
		ParentHash string `json:"parent_hash"`
		ParentName string `json:"parent_name"`
	} `json:"objects"`
	Project struct {
		Pid string `json:"pid"`
	} `json:"project"`
	Release struct {
		Release    string `json:"release"`
		CommitHash string `json:"commit_hash"`
	} `json:"release"`
}

type objectForPost struct {
	CommitTime string `json:"commit_time"`
	FilePath   string `json:"file_path"`
	Hash       string `json:"hash"`
	Name       string `json:"name"`
	OldName    string `json:"old_name"`
	Owner      string `json:"owner"`
	ParentHash string `json:"parent_hash"`
	ParentName string `json:"parent_name"`
}
