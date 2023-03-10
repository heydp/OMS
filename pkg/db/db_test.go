package db

import (
	"testing"
)

func TestInit(t *testing.T) {

	err := InitTest()
	if err != nil {
		t.Fatal("Err in run_main func")
	}
}
