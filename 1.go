package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var user1, user2 User
var item1, item2, item3, item4, item5 Item

func main() {
	item1 = NewItem(1, "Car", 1000.312)
	item2 = NewItem(2, "Pivo", 99.09)
	item3 = NewItem(3, "Kniga", 124.12)
	item4 = NewItem(4, "Iphone", 1434.1)
	item5 = NewItem(5, "Manga", 312.00)
	user1 = User{ID: 1, User: "user1"}
	user2 = User{ID: 2, User: "user2"}
	/*user1.AddItemsToCart(item1)
	user1.AddItemsToCart(item2)
	user2.AddItemsToCart(item2)
	user1.AddItemsToCart(item3)
	user1.AddItemsToCart(item4)
	user2.AddItemsToCart(item2)
	user2.AddItemsToCart(item5)
	user2.AddItemsToCart(item4)
	user1.RemoveItemsFromCart(item2)
	user2.RemoveItemsFromCart(item2)
	user2.RemoveItemsFromCart(item4)
	user1.ShowUserCart()
	user2.ShowUserCart()
	user1.itemPrice()*/
	//fmt.Println(user1.String())
	//fmt.Println(user2.String())
	//fmt.Println(user1.User, user1.Cart, user1.CartSum)
	//fmt.Println(user2.User, user2.Cart)

	log.Println("server start")
	http.HandleFunc("/get_catalog", catalogHandler)
	http.HandleFunc("/get_cart", cartHandler)
	http.HandleFunc("/get_user_cart", userHandler)
	http.HandleFunc("/add_item_to_cart", addHandler)
	http.HandleFunc("/remove_item_from_cart", removeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("OSHIBKA BLYA", err)
	}

}

type User struct {
	ID      int     `json:"ID"`
	User    string  `json:"User"`
	Cart    []Item  `json:"Cart"`
	CartSum float64 `json:"CartSum"`
}
type Item struct {
	ID    int     `json:"ID"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userid")
	itemID := r.URL.Query().Get("itemID")

	if userID == "" || itemID == "" {
		http.Error(w, "404", http.StatusBadRequest)
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}
	var user *User
	if userIDInt == user1.ID {
		user = &user1
	} else if userIDInt == user2.ID {
		user = &user2
	} else {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	item, err := getItemID(itemIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	user.RemoveItemsFromCart(item)
	user.itemPrice()
	w.WriteHeader(http.StatusOK)
}
func addHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userid")
	itemID := r.URL.Query().Get("itemid")

	if userID == "" || itemID == "" {
		http.Error(w, "параметры не указаны", http.StatusBadRequest)
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}
	var user *User
	if userIDInt == user1.ID {
		user = &user1
	} else if userIDInt == user2.ID {
		user = &user2
	} else {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	item, err := getItemID(itemIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	user.AddItemsToCart(item)
	user.itemPrice()
	w.WriteHeader(http.StatusOK)

}

func catalogHandler(w http.ResponseWriter, r *http.Request) {
	items := []Item{item1, item2, item3, item4, item5}
	json.NewEncoder(w).Encode(items)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userid")
	if userID == "" {
		http.Error(w, "405", http.StatusBadRequest)
		return
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "неверное значение", http.StatusBadRequest)
		return
	}
	var user *User
	if userIDInt == user1.ID {
		user = &user1
	} else if userIDInt == user2.ID {
		user = &user2
	} else {
		http.Error(w, "такого пользователя нет", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
	/*params := strings.Split(r.URL.Path, "/")
	if len(params) != 3 {
		http.Error(w, "404", http.StatusBadRequest)
		return
	}
	userName := params[2]
	var user User
	if userName == user1.User {
		user = user1
	} else if userName == user2.User {
		user = user2
	} else {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)*/
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCart(w, r)
	case http.MethodPost:
		postCart(w, r)
	default:
		http.Error(w, "invalid", http.StatusMethodNotAllowed)
	}
}
func postCart(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user1.AddItemsToCart(item)
	user1.itemPrice()
	fmt.Fprintf(w, "add new item: '%s'", item.Name)
}
func getCart(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user1.Cart)
}

func (u *User) itemPrice() {
	sum := 0.0
	for _, item := range u.Cart {
		sum += item.Price
	}
	u.CartSum = sum

}

func (u *User) AddItemsToCart(item Item) {
	u.Cart = append(u.Cart, item)

}
func (u *User) ShowUserCart() {

}
func (u *User) RemoveItemsFromCart(item Item) {
	for i, v := range u.Cart {
		if v == item {
			copy(u.Cart[i:], u.Cart[i+1:])
			u.Cart = u.Cart[:len(u.Cart)-1]
		}
	}
	return
}
func (i *Item) String() string {
	return fmt.Sprintf("%s: %.2f$", i.Name, i.Price)
}

func NewItem(ID int, Name string, Price float64) Item {
	return Item{ID, Name, Price}
}
func getItemID(itemID int) (Item, error) {
	for _, item := range []Item{item1, item2, item3, item4, item5} {
		if item.ID == itemID {
			return item, nil
		}
	}
	return Item{}, fmt.Errorf("такого товара нет", itemID)
}
func (u *User) String() string {
	cart := make([]string, len(u.Cart))
	for i, item := range u.Cart {
		cart[i] = fmt.Sprintf("%s: %.2f", item.Name, item.Price)
	}
	return fmt.Sprintf("%s, %s, %.2f$", u.User, strings.Join(cart, " "), u.CartSum)
}
