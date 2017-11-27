package domain

import (
	"time"
)

var idTweet int

// Tweet - Interfaz
type Tweet interface {
	PrintableTweet() string
	GetID() int
	GetUser() string
	GetText() string
	GetDate() *time.Time
}

// TextTweet - Define la estructura del tweet
type TextTweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

// NewTextTweet - Crea un nuevo tweet ingresando el texto y el usuario
func NewTextTweet(usr string, txt string) *TextTweet {
	date := time.Now()
	tweet := TextTweet{
		generateIDTweet(),
		usr,
		txt,
		&date,
	}
	return &tweet
}

// ImageTweet - Define la estructura del tweet
type ImageTweet struct {
	ID    int
	User  string
	Text  string
	Image string
	Date  *time.Time
}

// NewImageTweet - Crea un nuevo tweet ingresando el texto y el usuario
func NewImageTweet(usr string, txt string, img string) *ImageTweet {
	date := time.Now()
	tweet := ImageTweet{
		generateIDTweet(),
		usr,
		txt,
		img,
		&date,
	}
	return &tweet
}

// generateID - Genera un ID a cada nuevo tweet
func generateIDTweet() int {
	idTweet++
	return idTweet
}

// String -
func (t *TextTweet) String() string {
	return t.PrintableTweet()
}

// PrintableTweet -
func (t *TextTweet) PrintableTweet() string {
	return "@" + t.User + ": " + t.Text
}

// PrintableTweet -
func (t *ImageTweet) PrintableTweet() string {
	return "@" + t.User + ": " + t.Text + " " + t.Image
}

// GetUser -
func (t *TextTweet) GetUser() string {
	return t.User
}

// GetUser -
func (t *ImageTweet) GetUser() string {
	return t.User
}

// GetText -
func (t *TextTweet) GetText() string {
	return t.Text
}

// GetText -
func (t *ImageTweet) GetText() string {
	return t.Text
}

// GetID -
func (t *TextTweet) GetID() int {
	return t.ID
}

// GetID -
func (t *ImageTweet) GetID() int {
	return t.ID
}

// GetDate -
func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}

// GetDate -
func (t *ImageTweet) GetDate() *time.Time {
	return t.Date
}
