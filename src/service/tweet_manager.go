package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweet *domain.Tweet

// PublishTweet - Publicar tweet
func PublishTweet(twt *domain.Tweet) error {
	if twt.User == "" {
		return fmt.Errorf("user is required")
	} else if twt.Text == "" {
		return fmt.Errorf("text is required")
	} else if len(twt.Text) > 140 {
		return fmt.Errorf("tweet must be less than 140 chars")
	}
	tweet = twt
	return nil
}

// GetTweet - Devuelve tweet
func GetTweet() *domain.Tweet {
	return tweet
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func CleanTweet() {
	tweet = nil
}
