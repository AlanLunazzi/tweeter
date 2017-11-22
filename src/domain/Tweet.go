package domain

import (
	"time"
)

// Tweet - Define la estructura del tweet
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

// NewTweet - Crea un nuevo tweet ingresando el texto y el usuario
func NewTweet(usr string, txt string) *Tweet {
	date := time.Now()
	tweet := Tweet{
		usr,
		txt,
		&date,
	}
	return &tweet
}
