package daemon

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/ablades/prefix"
	"github.com/ablades/relevant/db"
	"github.com/ablades/relevant/models"
	"github.com/spf13/viper"
)

func parseSubPosts(sub *prefix.Tree, noteQueue chan models.Notification, waitGroup *sync.WaitGroup) {
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
	defer response.Body.Close()

	postContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

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
		fmt.Printf("Title: %s \n", post.Data.Title)
		fmt.Printf("PermaLink: %s \n", post.Data.Permalink)
		reg := regexp.MustCompile(`\w+`)
		parsedPost := reg.FindAllString(post.Data.Selftext+post.Data.Title, -1)
		fmt.Println(parsedPost)
		for _, word := range parsedPost {
			fmt.Println(word)
			users := sub.Contains(word)

			if len(users) > 0 {
				for _, user := range users {
					fmt.Printf("\033[32m Added Notification to channel for \033[0m user: \033[34m %s \033[0m  with word: \033[35m %s \033[0m \n", user, word)
					noteQueue <- models.Notification{
						Name:    user,
						Message: fmt.Sprintf("Post \033[34m %s 033[0m contains word \033[35m %s \033[0m \n Comment: \033[37m %s \033[0m \n", post.Data.Permalink, word, post.Data.Selftext),
					}
				}
			}

		}
	}
}

//Build Message from markdown template
func toMarkdown(masterMap map[string][]string) {
	f, _ := os.Create("file.md")
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	for key, value := range masterMap {
		err := t.Execute(f, models.MessageNote{User: key, Content: value})
		if err != nil {
			panic(err)
		}
	}
}

// Run Daemon
func Run() {

	//Connect to db
	db := db.Connect()
	subStore := models.GetSubRedditStore(db)
	subTrees := subStore.GetAllSubRedditTrees()

	var waitGroup sync.WaitGroup
	//Map Notifications to users
	noteMap := make(map[string][]string)
	noteQueue := make(chan models.Notification)

	for _, sub := range subTrees {
		fmt.Printf("\033[33m Init worker \033[0m for : \033[96m %s \033[0m \n", sub.Name)
		waitGroup.Add(1)
		go parseSubPosts(sub, noteQueue, &waitGroup)
	}

	//Read from channel until it's closed
	go func() {
		for note := range noteQueue {
			noteMap[note.Name] = append(noteMap[note.Name], note.Message)
		}
	}()

	waitGroup.Wait()
	close(noteQueue)
}
