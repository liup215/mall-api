package time

import (
	"testing"
)

func TestTime(t *testing.T) {
	now := Now()
	v, err := (&now).GetBSON()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Error(v)
	}
}
