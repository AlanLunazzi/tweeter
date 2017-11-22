package service

import (
	"github.com/tweeter/src/domain"
)

var tweet *domain.Tweet

// PublishTweet - Publicar tweet
func PublishTweet(twt *domain.Tweet) {
	tweet = twt
}

// GetTweet - Devuelve tweet
func GetTweet() *domain.Tweet {
	return tweet
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func CleanTweet() {
	tweet = nil
}
