package main

//Notification for a given user
type Notification struct {
	Subreddit string
	UserName  string
	Keyword   string
	Post      string
	PostTitle string
	Link      string
}

//MessageNote contains content of a post
type MessageNote struct {
	User    string
	Content []string
}
