package whosbugAssigns

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

const HOST = "https://sngapm.qq.com/whosbug"
const secret = "3E5D4C94-A9FE-4690-BEF4-76C40EAE44AB"
const userId = "qapm"

func genToken() string {
	expireStamp := time.Now()
	randNum := rand.Intn(10000000)
	v := fmt.Sprintf("%s||%s%8d%s", userId, expireStamp, randNum, secret)
	md5Value := md5.Sum([]byte(v))

	raw := fmt.Sprintf("%s||%s%8d%s", userId, expireStamp, randNum, md5Value)
	return raw
}
