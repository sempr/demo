package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fail()
	}
}

func TestAdd2(t *testing.T) {
	if Add(3, 2) != 5 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Fail()
	}
}
