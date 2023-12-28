// services/twitter.go

package services

/*
import (
	"context"
	"log"

	"github.com/g8rswimmer/go-twitter/v2"
	"golang.org/x/oauth2"
)

// TwitterService represents the service for interacting with the Twitter API
type TwitterService struct {
	client *twitter.Client
}

// NewTwitterService creates a new instance of TwitterService
func NewTwitterService(apiKey, apiSecretKey, accessToken, accessTokenSecret string) *TwitterService {
	// Set up Twitter API client with OAuth1 authentication
	config := &oauth2.Config{
		ClientID:     apiKey,
		ClientSecret: apiSecretKey,
		Scopes:       []string{"tweet.read", "user.read"},
		Endpoint:     twitter.Endpoint,
	}

	token := &oauth2.Token{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		RefreshToken: accessTokenSecret,
	}

	httpClient := config.Client(context.Background(), token)

	client := twitter.NewClient(httpClient)

	return &TwitterService{client: client}
}

// GetFollowerCount gets the number of followers for a given Twitter screen name
func (ts *TwitterService) GetFollowerCount(screenName string) (int, error) {
	// Retrieve user information including follower count
	user, _, err := ts.client.Users.Show(context.Background(), &twitter.UserShowParams{
		ScreenName: screenName,
	})
	if err != nil {
		log.Println("Error getting user information:", err)
		return 0, err
	}

	return user.Data.PublicMetrics.FollowersCount, nil
}
*/
