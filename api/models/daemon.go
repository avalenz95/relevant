package models

//Notification for a given user
type Notification struct {
	Name    string
	Message string
}

//MessageNote contains content of a post
type MessageNote struct {
	User    string
	Content []string
}
