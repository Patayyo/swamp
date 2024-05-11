package store

type Store interface {
	GetItems() ([]Item, error)
	AddItem(item Item) error
	GetItemByID(id string) (*Item, error)
	UpdateItem(id string, updatedItem Item) error
	DeleteItem(id string) error
	CreateUser(user User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByID(userID string) (*User, error)
	AddItemToCart(userID string, itemID string) error
	RemoveItemFromCart(userID, itemID string) error
	GetCart(userID string) ([]Item, error)
	GetUsers() ([]User, error)
	DeleteUser(id string) error
}
