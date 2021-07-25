package whosbugAssigns

type ChangeLineNumberType struct {
	LineNumber int
	ChangeType string
}
type DiffParsedType struct {
	DiffFile          string
	DiffFilePath      string
	ChangeLineNumbers []ChangeLineNumberType
	Commit            string
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

type analysisInfo struct {
	Diff string
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
