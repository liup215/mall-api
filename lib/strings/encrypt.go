package strings

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}
