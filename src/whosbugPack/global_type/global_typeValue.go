package global_type

type ImportHelper interface {
	Analysis()
}

// DiffParsedType 解析后的diff信息
type DiffParsedType struct {
	CommitterEmail    string
	CommitTime        string
	CommitterName     string
	DiffFileName      string
	ChangeLineNumbers []ChangeLineType
	CommitHash        string
	DiffText          string
	OldLineCount      int
	NewLineCount      int
}

// ChangeLineType 存储单个改变行的信息
type ChangeLineType struct {
	LineNumber int
	ChangeType string
}

// ObjectInfoType Ready For New Changes
type ObjectInfoType struct {
	CommitHash          string   `json:"commit_hash"`
	Id                  string   `json:"id"`
	OldId               string   `json:"old_id"`
	FilePath            string   `json:"file_path"`
	OldLineCount        int      `json:"old_line_count"`
	NewLineCount        int      `json:"new_line_count"`
	ChangedOldLineCount int      `json:"changed_old_line_count"`
	ChangedNewLineCount int      `json:"changed_new_line_count"`
	Calling             []string `json:"calling"`
}

func (s *ObjectInfoType) Equals(b ObjectInfoType) bool {
	if s.CommitHash == b.CommitHash {
		if s.Id == b.Id {
			return true
		}
	}
	return false
}

// InputJson 存储从input.json读取到的配置信息
type InputJson struct {
	ProjectId         string   `json:"__PROJECT_ID"`
	ReleaseVersion    string   `json:"__RELEASE_VERSION"`
	RepoPath          string   `json:"__PROJECT_URL"`
	BranchName        string   `json:"__BRANCH_NAME"`
	LanguageSupport   []string `json:"__LAN_SUPPORT"`
	WebServerHost     string   `json:"__WEB_SRV_HOST"`
	WebServerUserName string   `json:"__WEB_SRV_USERNAME__"`
	WebserverKey      string   `json:"__WEB_SRV_KEY__"`
	CryptoKey         string   `json:"__ENCRYPT_SECRET__"`
}

// 全局变量，存储config信息
var Config InputJson

// 全局变量，记录起始路径
var WorkPath string

// 全局变量，记录本地最后commit hash
var LocalHashLatest string

// 全局变量，object传输的通道
var ObjectChan chan ObjectInfoType

//// 全局变量，大型object传输的通道
//var ObjectChanLarge chan ObjectInfoType

// SupportLans 语言的支持
var SupportLans = []string{".java"}

var LatestCommitHash string

// CommitInfoType 存储每一次commit的信息
type CommitInfoType struct {
	CommitHash     string `json:"hash"`
	CommitterEmail string `json:"email"`
	CommitAuthor   string `json:"author"`
	CommitTime     string `json:"time"`
}
