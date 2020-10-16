package handlers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/gommon/log"
)

// Reddit API handlers

//handle get requests to reddit api
func (h *Handler) getRequestBytes(endpoint string) []byte {

	resp, err := h.client.Get(endpoint)
	if err != nil {
		log.Error(err)
	}

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
		// Inital content has not been set
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
	json.Unmarshal(content, &userInfo)

	return userInfo.Name
}
