package models

import (
	"context"

	"github.com/ablades/prefix"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SubReddit db representation
type SubReddit struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	BannerUrl string `json:"banner,omitempty" bson:"banner,omitempty"`
	//Color  string
	Tree prefix.Tree `json:"tree,omitempty" bson:"tree,omitempty"`
}

// SubRedditStore for db
type SubRedditStore struct {
	coll *mongo.Collection
}

// GetSubRedditStore from database
func GetSubRedditStore(db *mongo.Database) *SubRedditStore {
	return &SubRedditStore{
		coll: db.Collection("Subreddits"),
	}
}

// GetAllSubRedditNames from DB
func (subStore *SubRedditStore) GetAllSubRedditNames() []SubReddit {
	results, err := subStore.coll.Find(
		context.Background(),
		bson.D{},
		options.Find().SetProjection(bson.D{{Key: "name", Value: 1}}),
	)
	if err != nil {
		log.Error(err)
	}

	subreddits := &[]SubReddit{}

	results.All(context.Background(), subreddits)

	return *subreddits
}

func (subStore *SubRedditStore) CreateSubReddit(id string, name string, bannerUrl string) bool {
	result := subStore.coll.FindOne(context.Background(), bson.M{"_id": id})
	if result.Err() != mongo.ErrNoDocuments {
		return false
	}
	// resp, err := http.Get(bannerUrl)
	// if err != nil {
	// 	log.Error(err)
	// }

	// // img, fmtName, err := image.Decode(resp.Body)
	// // fmt.Print(fmtName)

	subreddit := SubReddit{
		ID:        id,
		Name:      name,
		BannerUrl: bannerUrl,
		Tree:      prefix.NewTree(name),
	}

	subStore.coll.InsertOne(context.Background(), subreddit)
	return true
}

func (subStore *SubRedditStore) UpdateTreeKeywords(subName string, userName string, keyword string) {
	result := subStore.coll.FindOne(context.Background(), bson.M{"name": subName})

	subreddit := &SubReddit{}
	result.Decode(subreddit)
	subreddit.Tree.InsertKeyword(keyword, userName)

	subStore.coll.ReplaceOne(context.Background(), bson.M{"name": subName}, subreddit)
}

// // CreateSubReddit Object TODO: Change this to a bulk insert/write
// func (subStore *SubRedditStore) CreateSubReddits(subredditsMap map[string]string) {

// 	for name, id := range subredditsMap {
// 		result := subStore.coll.FindOne(context.Background(), bson.M{"_id": id})
// 		// Does not exist add to db
// 		if result.Err() == mongo.ErrNoDocuments {
// 			subreddit := SubReddit{
// 				ID:   id,
// 				Name: name,
// 				Tree: prefix.NewTree(name),
// 			}
// 		}

// 	}

// 	// Give it an id
// 	// Name
// 	// Fetch a banner
// 	//Store banner image
// 	//Create a new prefix tree

// }
