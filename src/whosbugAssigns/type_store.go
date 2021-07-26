package whosbugAssigns

var supportLans = []string{".java"}

type ChangeLineNumberType struct {
	LineNumber int
	ChangeType string
}
type DiffParsedType struct {
	DiffFile          string
	DiffFilePath      string
	ChangeLineNumbers []ChangeLineNumberType
	CommitId          string
	DiffContent       map[int]map[string]string
}
type CommitterInfoType struct {
	Name  string
	Email string
}
type CommitParsedType struct {
	CommitLeftIndex int
	Commit          string
	CommitTime      string
	CommitterInfo   CommitterInfoType
	CommitDiffs     []DiffParsedType
}

type ReleaseDiffType struct {
	CommitInfo     string
	Diff           string
	BranchName     string
	HeadCommitInfo string
}

type ObjectInfoType struct {
	Name    string
	Hash    string
	ParName string
	ParHash string
}

type ChangeMethodType struct {
	StartLine    int
	MethodName   string
	MasterObject string
}

type input_json struct {
	ProjectId       string   `json:"__PROJRCT_ID"`
	ReleaseVersion  string   `json:"__RELEASE_VERSION"`
	ProjectUrl      string   `json:"__PROJECT_URL"`
	BranchName      string   `json:"__BRANCH_NAME"`
	LanguageSupport []string `json:"__LAN_SUPPORT"`
	WebServerHost   string   `json:"__WEB_SRV_HOST"`
}
type innerConfig struct {
	username string
	password string
}

var innerConf innerConfig
var Config input_json

func init() {
	innerConf.username = "kevinmatt"
	innerConf.password = "heyuheng1.22.3"
}
