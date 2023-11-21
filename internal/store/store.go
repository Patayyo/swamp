package store

type Store interface {
	GetItems() ([]Item, error)
	AddItem(item Item) error
	GetItemByID(id string) (*Item, error)
	UpdateItem(id string, updatedItem Item) error
	DeleteItem(id string) error
}
