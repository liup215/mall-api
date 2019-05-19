package strings

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func NumRandomStr(l int) string {
	if l > 10 {
		l = 10
	}
	return fmt.Sprintf("%v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(math.Pow10(l))))
}

//生成随机字符串
func Random(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result = append(result, bytes[r.Intn(62)])
	}
	return string(result)
}
