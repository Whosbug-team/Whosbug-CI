package whosbugPack

var supportLans = []string{".java"}

type inputJson struct {
	ProjectId       string   `json:"__PROJRCT_ID"`
	ReleaseVersion  string   `json:"__RELEASE_VERSION"`
	RepoPath        string   `json:"__PROJECT_URL"`
	BranchName      string   `json:"__BRANCH_NAME"`
	LanguageSupport []string `json:"__LAN_SUPPORT"`
	WebServerHost   string   `json:"__WEB_SRV_HOST"`
}

type commitInfoType struct {
	startLineNum int
	commitHash   string
	commitEmail  string
	commitName   string
	commitTime   string
}

type changeLineType struct {
	lineNumber int
	changeType string
}

type diffParsedType struct {
	diffFileName      string
	diffFilePath      string
	changeLineNumbers []changeLineType
	commitHash        string
	diffContent       map[int]map[string]string
}

type ChangeMethodType struct {
	StartLine    int
	MethodName   string
	MasterObject string
}
