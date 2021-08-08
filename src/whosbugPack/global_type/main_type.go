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
}

// ChangeLineType 存储单个改变行的信息
type ChangeLineType struct {
	LineNumber int
	ChangeType string
}

// ObjectInfoType
type ObjectInfoType struct {
	CommitTime string `json:"commit_time"`
	FilePath   string `json:"file_path"`
	Hash       string `json:"hash"`
	Name       string `json:"name"`
	OldName    string `json:"old_name"`
	Owner      string `json:"owner"`
	ParentHash string `json:"parent_hash"`
	ParentName string `json:"parent_name"`
}

// inputJson 存储从input.json读取到的配置信息
type InputJson struct {
	ProjectId         string   `json:"__PROJECT_ID"`
	ReleaseVersion    string   `json:"__RELEASE_VERSION"`
	RepoPath          string   `json:"__PROJECT_URL"`
	BranchName        string   `json:"__BRANCH_NAME"`
	LanguageSupport   []string `json:"__LAN_SUPPORT"`
	WebServerHost     string   `json:"__WEB_SRV_HOST"`
	WebServerUserName string   `json:"__WEB_SRV_USERNAME"`
	WebserverPassWord string   `json:"__WEB_SRV_PASSWORD"`
	CryptoSecret      string   `json:"__ENCRYPT_SECRET"`
}

// 全局变量，存储config信息
var Config InputJson

// 全局变量，记录起始路径
var WorkPath string

// 全局变量，记录本地最后commit hash
var LocalHashLatest string

// 全局变量，object传输的通道
var ObjectChan chan ObjectInfoType

// 全局变量，大型object传输的通道
var ObjectChanLarge chan ObjectInfoType

// SupportLans 语言的支持
var SupportLans = []string{".java"}

// Secret 加密密钥
var Secret string

// CommitInfoType 存储每一次commit的信息
type CommitInfoType struct {
	CommitHash     string
	CommitterEmail string
	CommitterName  string
	CommitTime     string
}
