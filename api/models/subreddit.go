package models

import (
	"github.com/ablades/prefix"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SubReddit db representation
type SubReddit struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Banner primitive.Binary   `json:"banner,omitempty" bson:"banner,omitempty"`
	Tree   prefix.Tree        `json:"tree,omitempty" bson:"tree,omitempty"`
}

// SubRedditStore for db
type SubRedditStore struct {
	coll *mongo.Collection
}

// GetSubRedditStore from database
func GetSubRedditStore(db *mongo.Database) *UserStore {
	return &UserStore{
		coll: db.Collection("Subreddits"),
	}
}

//fetchBanner from subreddit TODO Add other sources
func fetchBanner(subreddit SubReddit) {
	// 	url := fmt.Sprintf("https://www.reddit.com/%s/about.json", subname)
	// 	request, err := http.NewRequest("GET", url, nil)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	// 	content := sendRequest(request)
	// 	var as = aboutSubreddit{}
	// 	json.Unmarshal(content, &as)
	// 	fmt.Printf("Fetched Banner url: %s \n", as.Data.BannerImg)
	// 	return as.Data.BannerImg
	//
}

// CreateSubReddit Object
func (subStore *SubRedditStore) CreateSubReddit(subreddit SubReddit) {
	// Give it an id
	// Name
	// Fetch a banner
	//Store banner image
	//Create a new prefix tree

}
