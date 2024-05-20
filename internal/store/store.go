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
	AddItemToCart(string, string) error
	RemoveItemFromCart(string, string) error
	GetCart(string) ([]Item, error)
	CreateCart(string) error
	GetUsers() ([]User, error)
	DeleteUser(id string) error
	UpdateBalance(username string, amount float64) error
	GetBalance(username string) (float64, error)
}
