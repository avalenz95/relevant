package handlers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/gommon/log"
)

// Reddit API handlers

// Get all of a users subscribed subreddits
func getRedditUserSubs() {

}

// Get the users name from the api
func (h *Handler) getRedditUserName() string {

	resp, err := h.client.Get("https://oauth.reddit.com/api/v1/me")
	if err != nil {
		log.Error(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	userInfo := struct {
		Name string `json:"name"`
	}{}
	json.Unmarshal(content, &userInfo)

	return userInfo.Name
}
