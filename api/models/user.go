package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User holds basic user information
type User struct {
	ID   primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Name string              `json:"name,omitempty" bson:"name,omitempty"`
	Subs map[string][]string `json:"subs,omitempty" bson:"subs,omitempty"`
}
