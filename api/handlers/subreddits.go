package handlers

import (
	"net/http"

	"github.com/ablades/relevant/models"
	"github.com/labstack/echo/v4"
)

//UpdateKeywords in a user and in DB
func (h *Handler) UpdateKeywords(c echo.Context) (err error) {
	uStore := models.GetUserStore(h.db)
	subStore := models.GetSubRedditStore(h.db)
	userName := c.Param("name")
	subName := c.Param("subname")
	keyword := c.Param("keyword")

	subStore.UpdateTreeKeywords(subName, userName, keyword)
	updatedKeywords := uStore.UpdateUserKeywords(subName, userName, keyword)

	return c.JSON(http.StatusOK, updatedKeywords)
}

// // UpdateBanners in DB
// func (h *Handler) UpdateBanners(c echo.Context) (err error) {
// 	subStore := models.GetSubRedditStore(h.db)
// 	subreddits := subStore.GetAllSubRedditNames()

// 	// Process subreddits in batches
// 	// batchSize := 100
// 	// for i := 0; i < len(subreddits); i += batchSize {
// 	// 	//bounds catch
// 	// 	j := i + batchSize
// 	// 	if j > len(subreddits) {
// 	// 		j = len(subreddits)
// 	// 	}
// 	// 	h.fetchBanners(subreddits[i:j])
// 	// }

// 	//get a subs uid
// 	//use endpoint to get up to 100 subs at a time
// 	//https://www.reddit.com/api/info.json?id=t5_2qizd,t5_2u3r3,t5_2qh23
// 	// for _, subreddit := range subreddits {
// 	// 	subreddit.Name
// 	// }
// 	// //userName := h.getRedditUserName()
// 	// // Add list of subreddits to a user objects subs
// 	// subreddits := h.getRedditUserSubs()
// 	// subs := make(map[string][]string)
// 	// for _, subreddit := range subreddits {
// 	// 	subs[subreddit] = make([]string, 0)
// 	// }
// 	// // Create user
// 	// newUser := models.User{
// 	// 	Name: userName,
// 	// 	Subs: subs,
// 	// }

// 	// uStore.CreateUser(newUser)
// 	// //Insert user into db
// 	// return c.JSON(http.StatusCreated, user.Name)
// }
