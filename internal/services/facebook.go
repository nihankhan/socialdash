// facebook.go

package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// FacebookResponse represents the structure of the Facebook API response
type FacebookResponse struct {
	Friends FacebookFriends `json:"friends"`
}

// FacebookFriends represents the friends section in the Facebook API response
type FacebookFriends struct {
	Data    []FacebookFriend `json:"data"`
	Summary FacebookSummary  `json:"summary"`
}

// FacebookFriend represents a single friend in the Facebook API response
type FacebookFriend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// FacebookSummary represents the summary section in the Facebook API response
type FacebookSummary struct {
	TotalCount int `json:"total_count"`
}

// GetFacebookFriendCount fetches the friend count from Facebook API
func GetFacebookFriendCount() (int, error) {
	accessToken := "EAAE488av7pgBOzJss9caKEMj54MCBH7U04j4nU16EOnJONPZBs2hO9vv5YtZAWi92mNBp3dB83ivSXpYVq1st0H0ZA7DuZCZCnKrgHY6tOERuQlHMg4N6OQKHBpzXoqNorxQNlWTyQoZAybdXigrCZBGYH5sH43cWaZCVV8WnAKqT3dgdDAlQ6fWjlFr"
	userID := "400709028970275"
	url := fmt.Sprintf("https://graph.facebook.com/v12.0/%s?fields=friends.summary(true)&access_token=%s", userID, accessToken)

	// Make a GET request to the Facebook API
	response, err := resty.New().R().Get(url)
	if err != nil {
		return 0, err
	}

	// Parse the response
	var fbResponse FacebookResponse
	if err := json.Unmarshal(response.Body(), &fbResponse); err != nil {
		return 0, err
	}

	fmt.Println("Facebook: ", fbResponse.Friends.Summary.TotalCount)

	return fbResponse.Friends.Summary.TotalCount, nil
}
