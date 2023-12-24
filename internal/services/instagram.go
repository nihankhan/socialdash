// instagram.go

package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type InstagramConfig struct {
	AccessToken string
	// Add any other required fields for Instagram configuration
}

type InstagramResponse struct {
	Followers InstagramFollowers `json:"data"`
}

type InstagramFollowers struct {
	Count int `json:"follower_count"`
}

// GetInstagramFollowerCount fetches the follower count from Instagram API
func GetInstagramFollowerCount() (int, error) {
	accessToken := "IGQWRNUVQ5YnlsTkZAFN2FqS1N1RkJrMmt3N3NQOTJscmRXaE1mVXV6SkdnUGdSYm1qNUlVaXh2dXJCSTRqV1I3NXNVMHcyTjNlSHlxeUxPY09IOW1mUUNibWtZAVldoc2hkbk5xa1NVNG5VR2o2NDdXeTU3QVhlaWcZD"
	url := fmt.Sprintf("https://graph.instagram.com/v12.0/me?fields=followers_count&access_token=%s", accessToken)

	// Make a GET request to the Instagram API
	response, err := resty.New().R().Get(url)
	if err != nil {
		return 0, err
	}

	// Parse the response
	var igResponse InstagramResponse
	if err := json.Unmarshal(response.Body(), &igResponse); err != nil {
		return 0, err
	}

	return igResponse.Followers.Count, nil
}
