package db

type DBs interface {
	Connect() bool
	Status() bool
	Close() bool
}
