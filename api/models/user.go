package models

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User holds basic user information
type User struct {
	ID   primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Name string              `json:"name,omitempty" bson:"name,omitempty"`
	Subs map[string][]string `json:"subs,omitempty" bson:"subs,omitempty"`
}

// UserStore representation of the user collection
type UserStore struct {
	coll *mongo.Collection
}

// GetUserStore from database
func GetUserStore(db *mongo.Database) *UserStore {
	return &UserStore{
		coll: db.Collection("Users"),
	}
}

// CreateUser object
func (uStore *UserStore) CreateUser(user User) primitive.ObjectID {
	inserted, err := uStore.coll.InsertOne(context.Background(), user)
	if err != nil {
		log.Error(err)
	}

	return inserted.InsertedID.(primitive.ObjectID)
}

// GetUserByID from db
func (uStore *UserStore) GetUserByID(id primitive.ObjectID) *User {

	result := uStore.coll.FindOne(context.Background(), id)
	// Unmarshall into user object
	user := &User{}
	result.Decode(user)

	return user
}
