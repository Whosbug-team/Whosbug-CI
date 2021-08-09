package commit_diffpack

import (
	"fmt"
	"testing"
)

func TestQuatoToNum(t *testing.T) {
	text := "-1,20"
	res := QuatoToNum(text[1:])
	fmt.Println(res)
}
