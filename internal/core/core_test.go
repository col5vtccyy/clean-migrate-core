package core

import "testing"

func TestRun(t *testing.T) {
	if Run("a") != "a" {
		t.Fail()
	}
}
