package strings

import (
	"testing"
)

func TestNumRndomStr(t *testing.T) {
	str := NumRandomStr(100)
	t.Error(str)
}
