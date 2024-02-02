package test_go_mod

import (
	"testing"
)

func TestRemoteFunc(t *testing.T) {
	b := Test_go_mod()

	if b != true {
		t.Errorf("FAILED")
	}
}
