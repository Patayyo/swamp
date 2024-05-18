package store

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	Cart     []Item             `json:"cart,omitempty" bson:"cart,omitempty"`
	Role     string             `json:"role,omitempty" bson:"role,omitempty"`
	Balance  float64            `bson:"balance"`
}

type Item struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"Name" bson:"name,omitempty"`
	Price float64            `json:"Price" bson:"price,omitempty"`
	Type  string             `jsin:"Type" bson:"type,omitempty"`
}

type NewItemInput struct {
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}
