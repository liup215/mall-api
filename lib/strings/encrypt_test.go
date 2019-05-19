package strings

import (
	"testing"
)

func TestPhpmd5(t *testing.T) {
	pwd := "Lp19901108"
	salt := "vVvq77U7H4D7V07F"

	dis := "ded752c17241230480d00860c9094231"
	md5 := Phpmd5(pwd + salt)

	t.Error(md5, dis, md5 == dis)

}
