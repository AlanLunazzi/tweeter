package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweet *domain.Tweet

// PublishTweet - Publicar tweet
func PublishTweet(twt *domain.Tweet) error {
	var err error
	if twt.User == "" {
		err = fmt.Errorf("user is required")
	} else {
		tweet = twt
	}
	return err
}

// GetTweet - Devuelve tweet
func GetTweet() *domain.Tweet {
	return tweet
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func CleanTweet() {
	tweet = nil
}
