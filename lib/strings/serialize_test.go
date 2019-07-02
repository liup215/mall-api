package strings

import (
	"fmt"
	"testing"
)

func TestUnSerialize(t *testing.T) {

	type A struct {
		A string
		B string
	}

	a := A{
		A: "haha",
		B: "heihei",
	}

	d, err := PhpSerialize(a)
	fmt.Println(d)
	if err != nil {
		t.Error(err.Error())
	}

	var b A
	err = PhpUnserialize([]byte(d), &b)
	if err != nil {
		t.Error(err.Error())
	}

	t.Error("成功", b)

}
