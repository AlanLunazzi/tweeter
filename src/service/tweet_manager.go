package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

//var tweets []*domain.Tweet
var tweets map[int]*domain.Tweet
var tweetsByUser map[string][]int
var followersUser map[string][]string
var lastid int

// InitializeService - Asigna espacio en memoria a tweets
func InitializeService() {
	tweets = make(map[int]*domain.Tweet)
	tweetsByUser = make(map[string][]int)
	followersUser = make(map[string][]string)
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
	lastid = twt.ID
	tweets[twt.ID] = twt
	//elem, ok := tweetsByUser[twt.User]
	//if ok {
	tweetsByUser[twt.User] = append(tweetsByUser[twt.User], twt.ID)
	/*} else {
		tweetsByUser[twt.User] = make([]int, 0)
		tweetsByUser[twt.User] = append(tweetsByUser[twt.User], twt.ID)
	}*/
	return twt.ID, nil
}

// GetTweetByID - Devuelve tweet segun su id
func GetTweetByID(id int) *domain.Tweet {
	elem, ok := tweets[id]
	if ok {
		return elem
	}
	return nil
}

// GetTweetsByUser - Devuelve tweets segun su usuario
func GetTweetsByUser(user string) []*domain.Tweet {
	elem, ok := tweetsByUser[user]
	var userTweets []*domain.Tweet
	if ok {
		for _, id := range elem {
			userTweets = append(userTweets, tweets[id])
		}
		return userTweets
	}
	return nil
}

// GetTweet - Devuelve tweet
func GetTweet() *domain.Tweet {
	return tweets[lastid]
}

// GetTweets - Devuelve tweet
func GetTweets() []*domain.Tweet {
	var aux []*domain.Tweet
	for _, tweet := range tweets {
		aux = append(aux, tweet)
	}
	return aux
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func CleanTweet() {
	length := len(tweetsByUser[tweets[lastid].User])
	if length == 1 {
		tweetsByUser[tweets[lastid].User] = nil
	} else {
		tweetsByUser[tweets[lastid].User] = tweetsByUser[tweets[lastid].User][:(length - 1)]
	}
	delete(tweets, lastid)
	lastid--
}

// CountTweetsByUser - Contar tweets por usuario
func CountTweetsByUser(user string) int {
	return len(tweetsByUser[user])
}

// Follow - Agrega a un usuario otro usuario para seguirlo
func Follow(user string, followed string) {
	followersUser[user] = append(followersUser[user], followed)
}

// GetTimeline -
func GetTimeline(user string) []*domain.Tweet {
	var timeline []*domain.Tweet
	timeline = append(timeline, GetTweetsByUser(user)...)
	for _, usr := range followersUser[user] {
		timeline = append(timeline, GetTweetsByUser(usr)...)
	}
	return timeline
}
