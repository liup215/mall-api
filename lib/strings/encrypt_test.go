package strings

import (
	"testing"
)

func TestPhpmd5(t *testing.T) {
	pwd := "920222"
	salt := "tbqMzBhQNQsPy2Xu16tA"

	// dis := "ded752c17241230480d00860c9094231"
	dis := "77452275adc8c57c7d8fb893aa2a67d5"
	md5 := Md5(pwd + salt)

	t.Error(md5, dis, md5 == dis)

}
