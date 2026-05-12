package main

import (
	"day2/idsCheck"
	"testing"
)

func TestIsInvalidsIds(t *testing.T) {
	if idsCheck.IsInvalidId(11) == false {
		t.Errorf("11 is invalid")
	}
}

func TestIsValidIds(t *testing.T) {
	if idsCheck.IsInvalidId(12) != false {
		t.Errorf("12 is a valid id")
	}
}

func TestSomeInvalids(t *testing.T) {
	soma := idsCheck.SomeIds([]int{11, 1212, 3333})

	if soma != 4556 {
		t.Errorf("Esperado 4556, obteve %d", soma)
	}
}
