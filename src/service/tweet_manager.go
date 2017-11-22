package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweets []*domain.Tweet

// InitializeService - Asigna espacio en memoria a tweets
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
}

// PublishTweet - Publicar tweet
func PublishTweet(twt *domain.Tweet) (int, error) {
	if twt.User == "" {
		return 0, fmt.Errorf("user is required")
	} else if twt.Text == "" {
		return 0, fmt.Errorf("text is required")
	} else if len(twt.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	tweets = append(tweets, twt)
	return twt.ID, nil
}

// GetTweetByID - Devuelve tweet segun su id
func GetTweetByID(id int) *domain.Tweet {
	if len(tweets) > 0 {
		for i := 0; i < len(tweets); i++ {
			if tweets[i].ID == id {
				return tweets[i]
			}
		}
	}
	return nil
}

// GetTweet - Devuelve tweet
func GetTweet() *domain.Tweet {
	if len(tweets) > 0 {
		return tweets[len(tweets)-1]
	}
	return nil
}

// GetTweets - Devuelve tweet
func GetTweets() []*domain.Tweet {
	return tweets
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func CleanTweet() {
	if len(tweets) > 1 {
		tweets = tweets[0 : len(tweets)-1]
	} else {
		tweets = nil
	}
}
