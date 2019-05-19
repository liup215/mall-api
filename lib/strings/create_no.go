package strings

import (
	"time"
)

func CreateNo(prefix string) string {
	ts := time.Now().Format("200601021504")
	random := NumRandomStr(8)
	return prefix + ts + random
}
