package badger

import (
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
)

type BadgerDB struct {
	Instance badger.DB
}

func Discover(path string) ([]string, bool, error) {
	var exists bool
	dir, err := os.ReadDir(path)
	if err != nil {
		exists = false
		return nil, exists, err
	}
	names := make([]string, 0)
	for _, di := range dir {
		if !di.IsDir() &&
			filepath.Ext(filepath.Join(path, di.Name())) == ".db" {
			names = append(names, di.Name())
		}
	}
	if len(names) > 0 {
		exists = true
	}
	return names, exists, err
}

func NewOrExist(path string) *BadgerDB {
	return nil
}

func (bg *BadgerDB) Connect() {
	// Get the config file (yaml)
	//yaml.Unmarshal()
	badger.Open(badger.DefaultOptions(""))
}
