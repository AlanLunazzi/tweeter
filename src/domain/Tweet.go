package domain

import (
	"time"
)

var id int

// Tweet - Define la estructura del tweet
type Tweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

// NewTweet - Crea un nuevo tweet ingresando el texto y el usuario
func NewTweet(usr string, txt string) *Tweet {
	date := time.Now()
	tweet := Tweet{
		generateID(),
		usr,
		txt,
		&date,
	}
	return &tweet
}

// generateID - Genera un ID a cada nuevo tweet
func generateID() int {
	id++
	return id
}
