package whosbugAssigns

import "crypto/sha256"

func hashCode64(pid string, objectName string, filePath string) string {
	text := pid + objectName + filePath
	return string(sha256.New().Sum([]byte(text)))
}
