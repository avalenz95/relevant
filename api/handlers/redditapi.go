package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ablades/relevant/api/models"
	"github.com/labstack/gommon/log"
)

// Reddit API handlers

//handle get requests to reddit api
func (h *Handler) getRequestBytes(endpoint string) []byte {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Error(err)
	}

	req.Header.Add("User-Agent", os.Getenv("REDDIT_USER_AGENT"))

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
func (h *Handler) getRedditUserSubs() map[string]string {
	subStore := models.GetSubRedditStore(h.db)
	subredditsMap := make(map[string]string)

	// Subreddit JSON Data struct
	subredditJSON := struct {
		Data struct {
			Children []struct {
				Data struct {
					DisplayName string `json:"display_name"`
					//Subscribers         int    `json:"subscribers"`
					//Name                string `json:"name"`
					ID                    string `json:"id"`
					BannerImg             string `json:"banner_img"`
					BannerBackgroundColor string `json:"banner_background_color"`
					//DisplayNamePrefixed string `json:"display_name_prefixed"`
					//Description         string `json:"description"`
					//URL                 string `json:"url"`
				} `json:"data"`
			} `json:"children"`
			After string `json:"after"`
		} `json:"data"`
	}{}

	firstRequest := true
	for subredditJSON.Data.After == "" {

		var content []byte
		// Initial content has not been set
		if firstRequest == true {
			content = h.getRequestBytes("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100")
			json.Unmarshal(content, &subredditJSON)
			firstRequest = false
		} else {
			// Pagination - Use After for subsequent requests
			content = h.getRequestBytes("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100&after=" + subredditJSON.Data.After)
			json.Unmarshal(content, &subredditJSON)
		}
		// Add subs name and id to list
		for _, item := range subredditJSON.Data.Children {
			subredditsMap[item.Data.DisplayName] = item.Data.ID
			fmt.Println("--")
			fmt.Println(item.Data.DisplayName)
			subStore.CreateSubReddit(item.Data.ID, item.Data.DisplayName, item.Data.BannerImg)
		}

		fmt.Println("here")
	}

	return subredditsMap
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

// // UpdateBanners for subreddits
// func (h *Handler) fetchBanners(subreddits map[string]string) {
// 	u := &url.URL{
// 		Scheme: "https",
// 		Host:   "reddit.com",
// 		Path:   "api/info.json",
// 	}
// 	// Add subreddit id to the query
// 	for _, id := range subreddits {
// 		u.Query().Add("id", id)
// 	}
// 	url := u.Query().Encode()

// 	h.getRequestBytes(url)

// }

// func (h *Handler) fetchBanner(id string) {
// 	u := &url.URL{
// 		Scheme: "https",
// 		Host:   "reddit.com",
// 		Path:   "api/info.json",
// 	}

// 	u.Query().Add("id", id)

// 	url := u.Query().Encode()

// 	h.getRequestBytes(url)

//}
