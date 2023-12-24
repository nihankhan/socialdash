// twitter.go

package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	consumerKey    = "wgKgfm0vFL8mW6UKCIagz2RPO"
	consumerSecret = "stL2DR9AurgcmvsAu93JhNQ7TIYpNVedpXFqCzxQsPzbdAfVpM"
	accessToken    = "1662842349406539776-VhvK0na2wT9UGFdqqFMyQp7ALUofDx"
	accessSecret   = "GCNBG7IPiuNRZ7fGuX1JucI9zTRdi2YHklMZDrkVwcgB1"
	screenName     = "nihan3301"
)

// TwitterResponse represents the structure of the Twitter API response
type TwitterResponse struct {
	Data TwitterData `json:"data"`
}

// TwitterData represents the data section in the Twitter API response
type TwitterData struct {
	PublicMetrics TwitterPublicMetrics `json:"public_metrics"`
}

// TwitterPublicMetrics represents the public_metrics section in the Twitter API response
type TwitterPublicMetrics struct {
	FollowersCount int `json:"followers_count"`
}

// GetTwitterFollowerCount fetches the follower count from Twitter API
// GetTwitterFollowerCount fetches the follower count from Twitter API
func GetTwitterFollowerCount() (int, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/by/username/%s?user.fields=public_metrics", screenName)

	// Make a GET request to the Twitter API
	response, err := resty.New().R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		Get(url)

	if err != nil {
		return 0, err
	}

	// Check for errors in the response
	if response.StatusCode() != 200 {
		return 0, nil
	}

	// Parse the response
	var twitterResponse TwitterResponse
	if err := json.Unmarshal(response.Body(), &twitterResponse); err != nil {
		return 0, err
	}

	// Check for errors in the Twitter response
	if twitterResponse.Data.PublicMetrics.FollowersCount == 0 {
		return 0, nil //fmt.Errorf("Twitter API response does not contain valid follower count")
	}

	fmt.Println(twitterResponse.Data.PublicMetrics.FollowersCount)

	return twitterResponse.Data.PublicMetrics.FollowersCount, nil
}
