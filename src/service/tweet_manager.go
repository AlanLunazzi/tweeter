package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

// TweetManager - Te maneja los tweets papa
type TweetManager struct {
	//var tweets []*domain.Tweet
	tweets        map[int]*domain.Tweet
	tweetsByUser  map[string][]int
	followersUser map[string][]string
	lastid        int
}

// NewTweetManager - Asigna espacio en memoria a tweets
func NewTweetManager() *TweetManager {
	var tm TweetManager
	tm.lastid = 1
	tm.tweets = make(map[int]*domain.Tweet)
	tm.tweetsByUser = make(map[string][]int)
	tm.followersUser = make(map[string][]string)
	return &tm
}

// PublishTweet - Publicar tweet
func (tm TweetManager) PublishTweet(twt *domain.Tweet) (int, error) {
	if twt.User == "" {
		return 0, fmt.Errorf("user is required")
	} else if twt.Text == "" {
		return 0, fmt.Errorf("text is required")
	} else if len(twt.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	tm.lastid = twt.ID
	tm.tweets[twt.ID] = twt
	tm.tweetsByUser[twt.User] = append(tm.tweetsByUser[twt.User], twt.ID)
	return twt.ID, nil
}

// GetTweetByID - Devuelve tweet segun su id
func (tm TweetManager) GetTweetByID(id int) *domain.Tweet {
	elem, ok := tm.tweets[id]
	if ok {
		return elem
	}
	return nil
}

// GetTweetsByUser - Devuelve tweets segun su usuario
func (tm TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	elem, ok := tm.tweetsByUser[user]
	var userTweets []*domain.Tweet
	if ok {
		for _, id := range elem {
			userTweets = append(userTweets, tm.tweets[id])
		}
		return userTweets
	}
	return nil
}

// GetTweet - Devuelve tweet
func (tm TweetManager) GetTweet() *domain.Tweet {
	return tm.tweets[tm.lastid]
}

// GetTweets - Devuelve tweet
func (tm TweetManager) GetTweets() []*domain.Tweet {
	var aux []*domain.Tweet
	for _, tweet := range tm.tweets {
		aux = append(aux, tweet)
	}
	return aux
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func (tm TweetManager) CleanTweet() {
	length := len(tm.tweetsByUser[tm.tweets[tm.lastid].User])
	if length == 1 {
		tm.tweetsByUser[tm.tweets[tm.lastid].User] = nil
	} else {
		tm.tweetsByUser[tm.tweets[tm.lastid].User] = tm.tweetsByUser[tm.tweets[tm.lastid].User][:(length - 1)]
	}
	delete(tm.tweets, tm.lastid)
	tm.lastid--
}

// CountTweetsByUser - Contar tweets por usuario
func (tm TweetManager) CountTweetsByUser(user string) int {
	return len(tm.tweetsByUser[user])
}

// Follow - Agrega a un usuario otro usuario para seguirlo
func (tm TweetManager) Follow(user string, followed string) {
	tm.followersUser[user] = append(tm.followersUser[user], followed)
}

// GetTimeline -
func (tm TweetManager) GetTimeline(user string) []*domain.Tweet {
	var timeline []*domain.Tweet
	timeline = append(timeline, tm.GetTweetsByUser(user)...)
	for _, usr := range tm.followersUser[user] {
		timeline = append(timeline, tm.GetTweetsByUser(usr)...)
	}
	return timeline
}
