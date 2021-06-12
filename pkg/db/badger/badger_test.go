package badger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscover(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(filepath.Dir(wd))
	names, ext, err := Discover(filepath.Dir(wd))
	if err != nil {
		t.Error(err.Error())
	}
	if ext {
		t.Log(names)
		t.Log(len(names))
	}
}
