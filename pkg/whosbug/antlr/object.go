package antlr

// 全局变量，object传输的通道
var ObjectChan chan ObjectInfoType

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
//
//	@receiver s *ObjectInfoType
//	@param b ObjectInfoType
//	@return bool
//	@author: Kevineluo 2022-07-31 07:09:43
func (s *ObjectInfoType) Equals(b ObjectInfoType) bool {
	if s.CommitHash == b.CommitHash && s.ID == b.ID && s.FilePath == b.FilePath && s.StartLine == b.StartLine && s.EndLine == b.EndLine {
		return true
	}
	return false
}
