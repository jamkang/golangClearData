package rand_te

import (
	"math/rand"
	"time"
)

//获取随机数
func GetRand(n ...int) int {
	rand.Seed(time.Now().UnixNano())
	if len(n) == 0 {
		return rand.Int()
	} else {
		return rand.Intn(n[0])
	}
}
