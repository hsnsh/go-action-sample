package adt_test

import (
	"go-action-sample/adt"
	"testing"
)

func TestEmpty(t *testing.T) {
	s := adt.NewStack()
	if s.Empty() != true {
		t.Error("Stack was not empty")
	}
}

func TestNotEmpty(t *testing.T) {
	s := adt.NewStack()
	s.Add("Hasan")
	if s.Empty() == true {
		t.Error("Stack was empty")
	}
}
