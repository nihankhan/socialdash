// youtube.go

package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

// YouTubeResponse represents the structure of the YouTube API response
type YouTubeResponse struct {
	Items []YouTubeItem `json:"items"`
}

// YouTubeSubscribers represents the subscribers section in the YouTube API response
type YouTubeItem struct {
	Statistics YouTubeStatistics `json:"statistics"`
}

type YouTubeStatistics struct {
	SubscriberCount string `json:"subscriberCount"`
}

// GetYouTubeSubscriberCount fetches the subscriber count from the YouTube API
func GetYouTubeSubscriberCount() (int, error) {
	channelID := "UCNzS7lrzRcTKgMD8cdX8DUA"
	apiKey := "AIzaSyBK2IcociFMU1e7iJ41nW-drY1EipRHOdA"
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/channels?part=statistics&id=%s&key=%s", channelID, apiKey)

	// Make a GET request to the YouTube API
	response, err := resty.New().R().Get(url)
	if err != nil {
		return 0, err
	}

	// Parse the response
	var ytResponse YouTubeResponse
	if err := json.Unmarshal(response.Body(), &ytResponse); err != nil {
		return 0, err
	}

	//fmt.Println(ytResponse.Items.Statistics.SubscriberCount)

	if len(ytResponse.Items) > 0 {
		subscriberCount, err := strconv.Atoi(ytResponse.Items[0].Statistics.SubscriberCount)

		if err != nil {
			return 0, err
		}

		return subscriberCount, nil
	}

	return 0, nil
}

//https://www.googleapis.com/youtube/v3/channels?part=statistics&id=UCNzS7lrzRcTKgMD8cdX8DUA&key=AIzaSyBK2IcociFMU1e7iJ41nW-drY1EipRHOdA
