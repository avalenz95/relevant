// package models

// import "go.mongodb.org/mongo-driver/mongo"

// // Panoptic holds names and ids of all users and subs
// type PanopticView struct {
// 	Subs  PanopticSubsView
// 	Users PanopticSubsView
// }

// type PanopticUsersView struct {
// 	Users map[string]string
// }

// type PanopticSubsView struct {
// 	Subs map[string]string
// }

// // PanopticStore rep of all users and subs
// type PanopticStore struct {
// 	coll *mongo.Collection
// }

// func GetPanopticStore(db *mongo.Database) *PanopticStore {
// 	return &PanopticStore{
// 		coll: db.Collection("Panoptic"),
// 	}
// }
