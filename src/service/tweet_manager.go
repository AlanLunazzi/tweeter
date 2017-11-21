package service

var tweet string

// PublishTweet - Publicar tweet
func PublishTweet(tw string) {
	tweet = tw
}

// GetTweet - Devuelve tweet
func GetTweet() string {
	return tweet
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func CleanTweet() {
	tweet = ""
}
