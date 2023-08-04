package convert

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStringToInteger(t *testing.T) {
	strVal := "1337"
	intVar := 1337

	// strconv.Atoi()
	intVar1, err := strconv.Atoi(strVal)
	switch err {
	case nil:
		if intVar1 != intVar {
			t.Errorf("strconv.Atoi(%q) is %d, expected %d", strVal, intVar1, intVar)
		}
	default:
		t.Errorf("strconv.Atoi(%q) failed with: %s", strVal, err)
	}

	// strconv.ParseInt()
	intVar2, err := strconv.ParseInt(strVal, 0, 64)
	switch err {
	case nil:
		if intVar2 != int64(intVar) {
			t.Errorf("strconv.ParseInt(%q, 0, 64) is %d, expected %d", strVal, intVar2, intVar)
		}
	default:
		t.Errorf("strconv.ParseInt(%q, 0, 64) failed with: %s", strVal, err)
	}

	// fmt.Sscan()
	var intVar3 int
	ne := 1
	n, err := fmt.Sscan(strVal, &intVar3)
	switch err {
	case nil:
		if n != ne {
			t.Errorf("fmt.Sscan(%q, &intVar3): intVar3 is %d, expected %d", strVal, n, ne)
		}
		if intVar3 != intVar {
			t.Errorf("fmt.Sscan(%q, &intVar3): intVar3 is %d, expected %d", strVal, intVar3, intVar)
		}
	default:
		t.Errorf("fmt.Sscan(%q, &intVar3) failed with: %s", strVal, err)
	}
}
