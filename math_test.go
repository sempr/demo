package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Fail()
	}
}

func TestSub2(t *testing.T) {
	if Sub(1, 1) != 0 {
		t.Fail()
	}
}

func TestSub3(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Fail()
	}
}
