package models

import (
	"github.com/ablades/prefix"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SubReddit db representation
type SubReddit struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Subname string             `json:"subname,omitempty" bson:"subname,omitempty"`
	Banner  primitive.Binary   `json:"banner,omitempty" bson:"banner,omitempty"`
	Tree    prefix.Tree        `json:"tree,omitempty" bson:"tree,omitempty"`
}

// SubRedditStore for db
type SubRedditStore struct {
	coll *mongo.Collection
}

// GetSubRedditStore from database
func GetSubRedditStore(db *mongo.Database) *UserStore {
	return &UserStore{
		coll: db.Collection("subreddits"),
	}
}

// CreateSubReddit Object
func (subStore *SubRedditStore) CreateSubReddit(subreddit SubReddit) {

}
