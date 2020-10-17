package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// Reddit API handlers

//handle get requests to reddit api
func (h *Handler) getRequestBytes(endpoint string) []byte {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Error(err)
	}

	req.Header.Add("User-Agent", viper.GetString("reddit.agent"))

	resp, err := h.client.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	return content
}

// Get all of a users subscribed subreddits
func (h *Handler) getRedditUserSubs() []string {

	subreddits := make([]string, 0)
	// Subreddit JSON Data struct
	subreddit := struct {
		Data struct {
			Children []struct {
				Data struct {
					DisplayName string `json:"display_name"`
					//Subscribers         int    `json:"subscribers"`
					//Name                string `json:"name"`
					//ID                  string `json:"id"`
					//DisplayNamePrefixed string `json:"display_name_prefixed"`
					//Description         string `json:"description"`
					//URL                 string `json:"url"`
				} `json:"data"`
			} `json:"children"`
			After string `json:"after"`
		} `json:"data"`
	}{}

	var content []byte
	for {
		// Initial content has not been set
		if content == nil {
			content := h.getRequestBytes("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100")
			json.Unmarshal(content, &subreddit)
		} else {
			// Pagination - Use After for subsequent requests
			content := h.getRequestBytes("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100&after=" + subreddit.Data.After)
			json.Unmarshal(content, &subreddit)
		}
		// Add subs to list
		for _, item := range subreddit.Data.Children {
			subreddits = append(subreddits, item.Data.DisplayName)
		}
		// continue till no more results available
		if subreddit.Data.After == "" {
			break
		}
	}

	return subreddits
}

// Get the users name from the api
func (h *Handler) getRedditUserName() string {
	// User info struct
	userInfo := struct {
		Name string `json:"name"`
	}{}

	content := h.getRequestBytes("https://oauth.reddit.com/api/v1/me")
	fmt.Println(string(content))
	json.Unmarshal(content, &userInfo)
	fmt.Println(userInfo)
	return userInfo.Name
}
