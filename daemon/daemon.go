package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
	"text/template"

	"github.com/ablades/prefix"
	"github.com/ablades/relevant/api/db"
	"github.com/ablades/relevant/api/models"
	"github.com/spf13/viper"
)

func parseSubPosts(sub prefix.Tree, noteQueue chan Notification, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	defer fmt.Printf("\033[32m Done parsing: \033[35m %s! \033[0m \n", sub.Name)

	url := fmt.Sprintf("https://api.reddit.com/r/%s/new", sub.Name) // best temporarily for consistent input data

	//send a request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", viper.GetString("reddit.name")))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	postContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	response.Body.Close()

	fmt.Printf("VISITING: %s \n", url)

	posts := struct {
		Kind string `json:"kind"`
		Data struct {
			Modhash  string `json:"modhash"`
			Dist     int    `json:"dist"`
			Children []struct {
				Kind string `json:"kind"`
				Data struct {
					ApprovedAtUtc       interface{} `json:"approved_at_utc"`
					Subreddit           string      `json:"subreddit"`
					Selftext            string      `json:"selftext"`
					Title               string      `json:"title"`
					Name                string      `json:"name"`
					Ups                 int         `json:"ups"`
					TotalAwardsReceived int         `json:"total_awards_received"`
					Edited              bool        `json:"edited"`
					ContentCategories   interface{} `json:"content_categories"`
					Created             float64     `json:"created"`
					ViewCount           interface{} `json:"view_count"`
					Archived            bool        `json:"archived"`
					Score               int         `json:"score"`
					Over18              bool        `json:"over_18"`
					Spoiler             bool        `json:"spoiler"`
					Locked              bool        `json:"locked"`
					SubredditID         string      `json:"subreddit_id"`
					Author              string      `json:"author"`
					NumComments         int         `json:"num_comments"`
					Permalink           string      `json:"permalink"`
					URL                 string      `json:"url"`
					CreatedUtc          float64     `json:"created_utc"`
				} `json:"data"`
			} `json:"children"`
			After string `json:"after"`
		} `json:"data"`
	}{}

	//parse json subreddit struct
	json.Unmarshal(postContent, &posts)

	for _, post := range posts.Data.Children {
		reg := regexp.MustCompile(`\w+`)
		parsedPost := reg.FindAllString(post.Data.Selftext+post.Data.Title, -1)
		//fmt.Println(parsedPost)
		for _, word := range parsedPost {
			//fmt.Println(word)
			users := sub.Contains(word)

			if len(users) > 0 {
				for _, user := range users {
					fmt.Printf("\033[32m Added Notification to channel for \033[0m user: \033[34m %s \033[0m  with word: \033[35m %s \033[0m \n", user, word)
					fmt.Printf("PermaLink: %s \n", post.Data.Permalink)
					noteQueue <- Notification{
						Subreddit: post.Data.Subreddit,
						UserName:  user,
						PostTitle: post.Data.Title,
						Post:      post.Data.Selftext,
						Keyword:   word,
						Link:      post.Data.URL,
						// Timestamp: post.Data.CreatedUtc
						// post.Data.
					}
				}
			}

		}
	}
}

//Build Message from markdown template
func toMarkdown(masterMap map[string]map[string][]Notification) {
	// Create Template
	t := template.Must(template.New("template.tmpl").ParseFiles("./template.tmpl"))
	// Iter over users
	for user, subMap := range masterMap {
		// Create File for a user
		f, _ := os.Create(user + "output.md")
		// Execute template on map
		err := t.Execute(f, subMap)
		if err != nil {
			panic(err)
		}
	}
}

// Run Daemon
func main() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %s", err))
	}
	//Connect to db
	db := db.Connect()
	subStore := models.GetSubRedditStore(db)
	subreddits := subStore.GetAllSubReddits()

	var waitGroup sync.WaitGroup
	//Note Map  - Notification Map  map[user] -> map[subreddit][posts]
	noteMap := make(map[string]map[string][]Notification)
	noteQueue := make(chan Notification)

	for _, sub := range subreddits {
		fmt.Printf("\033[33m Init worker \033[0m for : \033[96m %s \033[0m \n", sub.Name)
		waitGroup.Add(1)
		go parseSubPosts(sub.Tree, noteQueue, &waitGroup)
	}

	//Read from channel until it's closed
	go func() {
		for note := range noteQueue {

			//Check for nested map create if doesn't exist
			subMap, ok := noteMap[note.UserName]
			if !ok {
				subMap = make(map[string][]Notification)
				noteMap[note.UserName] = subMap
			}
			subMap[note.Subreddit] = append(subMap[note.Subreddit], note)
		}
	}()

	waitGroup.Wait()
	close(noteQueue)

	toMarkdown(noteMap)
}
