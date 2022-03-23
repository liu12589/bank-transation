package util

import (
	"math/rand"
	"strings"
	"time"
)

/*
	1、生成一个随机数 int64
	2、生成一个字符串
*/

const alphabet = "abcdefghigklmnopqrstuvwxyz"

func init() {
	// 每次调用初始化不同的随机数种子
	rand.Seed(time.Now().Unix())
}

// GenerateInt 生成随机数
func GenerateInt(min, max int64) int64  {
	return min + rand.Int63n(max - min + 1)
}

func GenerateString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func GenerateOwner() string {
	return GenerateString(6)
}

func GenerateMoney() int64 {
	return GenerateInt(0,1000)
}

func GenerateCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}