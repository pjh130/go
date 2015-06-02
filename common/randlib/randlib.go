package randlib

import (
	"math/rand"
	"time"
)

//生成随机数字
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
