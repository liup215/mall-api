package strings

import (
	"fmt"
	"testing"
)

func TestPhpmd5(t *testing.T) {
	pwd := "123456"
	salt := "i0Qi16ZrzF"

	// dis := "46f00be953306350a5b31580f11dfc22ee58d166"
	sha1 := Sha1(fmt.Sprintf("%s-%s-%s", pwd, salt, "35eb5e04"))

	t.Error(sha1)

}
