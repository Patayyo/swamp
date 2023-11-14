package store

type Store interface {
	GetItems() ([]Item, error)
}
