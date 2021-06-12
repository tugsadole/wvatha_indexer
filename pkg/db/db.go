package db

type DBType int

const (
	badger DBType = 0
	bolt   DBType = 1
)

type DBs interface {
	Discover()
	Connect() bool
	Status() bool
	Close() bool
}
