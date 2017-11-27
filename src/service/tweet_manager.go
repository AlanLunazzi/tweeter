package service

import (
	"fmt"
	"strings"

	"github.com/tweeter/src/domain"
)

// TweetManager - Te maneja los tweets papa
type TweetManager struct {
	//var tweets []domain.Tweet
	tweets        map[int]domain.Tweet
	tweetsByUser  map[string][]int
	followersUser map[string][]string
	lastid        int
	userMessages  map[string][]*domain.DirectMessage
}

// NewTweetManager - Asigna espacio en memoria a tweets
func NewTweetManager() *TweetManager {
	var tm TweetManager
	tm.lastid = 1
	tm.tweets = make(map[int]domain.Tweet)
	tm.tweetsByUser = make(map[string][]int)
	tm.followersUser = make(map[string][]string)
	tm.userMessages = make(map[string][]*domain.DirectMessage)
	return &tm
}

// PublishTweet - Publicar tweet
func (tm TweetManager) PublishTweet(twt domain.Tweet) (int, error) {
	if twt.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	} else if twt.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	} else if len(twt.GetText()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	tm.lastid = twt.GetID()
	tm.tweets[twt.GetID()] = twt
	tm.tweetsByUser[twt.GetUser()] = append(tm.tweetsByUser[twt.GetUser()], twt.GetID())
	return twt.GetID(), nil
}

// GetTweetByID - Devuelve tweet segun su id
func (tm TweetManager) GetTweetByID(id int) domain.Tweet {
	elem, ok := tm.tweets[id]
	if ok {
		return elem
	}
	return nil
}

// GetTweetsByUser - Devuelve tweets segun su usuario
func (tm TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	elem, ok := tm.tweetsByUser[user]
	var userTweets []domain.Tweet
	if ok {
		for _, id := range elem {
			userTweets = append(userTweets, tm.tweets[id])
		}
		return userTweets
	}
	return nil
}

// GetTweet - Devuelve tweet
func (tm TweetManager) GetTweet() domain.Tweet {
	return tm.tweets[tm.lastid]
}

// GetTweets - Devuelve tweet
func (tm TweetManager) GetTweets() []domain.Tweet {
	var aux []domain.Tweet
	for _, tweet := range tm.tweets {
		aux = append(aux, tweet)
	}
	return aux
}

// CleanTweet - Borra el ultimo tweet reemplazandolo por un texto vacio
func (tm TweetManager) CleanTweet() {
	length := len(tm.tweetsByUser[tm.tweets[tm.lastid].GetUser()])
	if length == 1 {
		tm.tweetsByUser[tm.tweets[tm.lastid].GetUser()] = nil
	} else {
		tm.tweetsByUser[tm.tweets[tm.lastid].GetUser()] = tm.tweetsByUser[tm.tweets[tm.lastid].GetUser()][:(length - 1)]
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
func (tm TweetManager) GetTimeline(user string) []domain.Tweet {
	var timeline []domain.Tweet
	timeline = append(timeline, tm.GetTweetsByUser(user)...)
	for _, usr := range tm.followersUser[user] {
		timeline = append(timeline, tm.GetTweetsByUser(usr)...)
	}
	return timeline
}

// GetTrendingTopics - Devuelve un arrays de las dos strings mas repetidas
func (tm TweetManager) GetTrendingTopics() [2]string {
	var allTweets string
	var primero, segundo int
	var keyPrimero, keySegundo string
	for _, tweet := range tm.tweets {
		allTweets = allTweets + " " + tweet.GetText()
	}
	allTweets = strings.ToLower(allTweets)
	mapa := wordCount(allTweets)
	for key, elem := range mapa {
		if elem > primero {
			segundo = primero
			primero = elem
			keySegundo = keyPrimero
			keyPrimero = key
		} else if elem > segundo {
			segundo = elem
			keySegundo = key
		}
	}
	return [2]string{keyPrimero, keySegundo}
}

//SendDirectMessage - Envia mensaje directo a un usuario
func (tm TweetManager) SendDirectMessage(user string, userTo string, text string) error {
	if user == "" {
		return fmt.Errorf("user is required")
	} else if userTo == "" {
		return fmt.Errorf("userTo is required")
	} else if text == "" {
		return fmt.Errorf("message is required")
	}

	direct := domain.CreateMessage(user, userTo, text)
	tm.userMessages[userTo] = append(tm.userMessages[userTo], direct)
	return nil
}

//ReadDirectMessage - Devuelve el dm del usuario que se corresponde al id
func (tm TweetManager) ReadDirectMessage(userTo string, id int) *domain.DirectMessage {
	length := len(tm.userMessages[userTo])
	if length > 0 {
		for _, dm := range tm.userMessages[userTo] {
			if dm.ID == id {
				dm.Read = true
				return dm
			}
		}
	}
	return nil
}

//GetUnreadDm - Devuelve los dm no leidos por el usuario
func (tm TweetManager) GetUnreadDm(userTo string) []*domain.DirectMessage {
	var noLeidos = make([]*domain.DirectMessage, 0)
	length := len(tm.userMessages[userTo])
	if length > 0 {
		for _, dm := range tm.userMessages[userTo] {
			if dm.Read == false {
				noLeidos = append(noLeidos, dm)
				return noLeidos
			}
		}
	}
	return nil
}

// GetAllDirectMessages - Returns a MAP with all the messages
func (tm TweetManager) GetAllDirectMessages(userTo string) []*domain.DirectMessage {
	return tm.userMessages[userTo]
}

func wordCount(s string) map[string]int {
	var m = make(map[string]int)
	wordsArray := strings.Split(s, " ")
	for _, word := range wordsArray {
		elem, ok := m[word]
		if ok {
			m[word] = elem + 1
		} else {
			m[word] = 1
		}
	}
	return m
}
