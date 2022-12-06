package main

import (
	"testing"
)

func TestCheckDup(t *testing.T) {
	if HasDuplicates("mjqj") {
		t.Log("It has duplicates")
	} else {
		t.Error("The string had dupplicates")
	}
}
