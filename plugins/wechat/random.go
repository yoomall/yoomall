package wechat

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes 随机字符串
func randStringBytes(n int) string {
	// 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
