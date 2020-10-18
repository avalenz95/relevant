package models

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
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

// UpdateUserSubs add new subscriptions from reddit into DB
func (uStore *UserStore) UpdateUserSubs(name string, subList map[string]string) bool {
	result := uStore.coll.FindOne(context.Background(), bson.M{"name": name})
	if result.Err() == mongo.ErrNoDocuments {
		log.Error(result.Err())
		return false
	}

	user := &User{}
	result.Decode(user)

	// Make a new map, case handles user unsubscribing and subscribing to various subreddits
	subMap := make(map[string][]string)
	// Loop over subs
	for subName, _ := range subList {
		// Check to see if subName is in old map
		keywords, contains := user.Subs[subName]
		if contains {
			subMap[subName] = keywords
		} else {
			subMap[subName] = make([]string, 0)
		}
	}
	//Reassign subs
	user.Subs = subMap

	//Replace user object TODO: Determine if it's theres a performance boon by just updating subs instead of re-adding
	_, err := uStore.coll.ReplaceOne(context.Background(), bson.M{"name": name}, user)
	if err != nil {
		log.Error(err)
	}

	return true
}

// CreateUser object
func (uStore *UserStore) CreateUser(user User) primitive.ObjectID {
	inserted, err := uStore.coll.InsertOne(context.Background(), user)
	if err != nil {
		log.Error(err)
	}

	return inserted.InsertedID.(primitive.ObjectID)
}

// GetUserByName from db
func (uStore *UserStore) GetUserByName(name string) *User {

	result := uStore.coll.FindOne(context.Background(), bson.M{"name": name})
	// Unmarshall into user object
	if result.Err() == mongo.ErrNoDocuments {
		return nil
	}

	// Unmarshall into user object
	user := &User{}
	result.Decode(user)

	return user
}

// DeleteUserByName from DB
func (uStore *UserStore) DeleteUserByName(name string) bool {
	_, err := uStore.coll.DeleteOne(context.Background(), bson.M{"name": name})
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (uStore *UserStore) UpdateUserKeywords(subName string, userName string, keyword string) []string {
	result := uStore.coll.FindOne(context.Background(), bson.M{"name": userName})

	// Unmarshall into user object
	user := &User{}
	result.Decode(user)

	user.Subs[subName] = append(user.Subs[subName], keyword)
	uStore.coll.ReplaceOne(context.Background(), bson.M{"name": userName}, user)

	return user.Subs[subName]
}
