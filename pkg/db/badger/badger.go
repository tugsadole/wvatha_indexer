package badger

import (
	"github.com/dgraph-io/badger/v3"
)

type BadgerDB struct {
	Instance badger.DB
}

func (bg *BadgerDB) Connect() {
	// Get the config file (yaml)
	//yaml.Unmarshal()
	badger.Open(badger.DefaultOptions(""))
}
