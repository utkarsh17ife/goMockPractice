package store

type DbStore interface {
	Insert(data interface{}) bool
}
