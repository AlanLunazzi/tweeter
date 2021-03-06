package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func main() {
	shell := ishell.New()
	tweetManager := service.NewTweetManager()

	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands \n")
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			var err error
			var id int
			var tweet domain.Tweet
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			usr := c.ReadLine()
			c.Print("Write your tweet: ")
			txt := c.ReadLine()
			c.Println("Enter the type of tweet you want to publish: (t - text i - image)")
			tt := c.ReadLine()
			if tt == "t" {
				tweet = domain.NewTextTweet(usr, txt)
			} else if tt == "i" {
				c.Print("Enter your url: ")
				url := c.ReadLine()
				tweet = domain.NewImageTweet(usr, txt, url)
			} else {
				return
			}
			id, err = tweetManager.PublishTweet(tweet)
			if err != nil {
				c.Println("Error->", err)
			} else {
				c.Println("Tweet Sent with id ", id)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := tweetManager.GetTweet()
			if tweet == nil {
				c.Println("No hay ultimo tweet")
			} else {
				c.Println(tweet)
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetByID",
		Help: "Shows a tweet by its ID",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Insert ID: ")
			var id int
			var err error
			id, err = strconv.Atoi(c.ReadLine())
			if err == nil {
				tweet := tweetManager.GetTweetByID(id)
				if tweet == nil {
					c.Println("No tweet with matching ID")
				} else {
					c.Println(tweet)
				}
			} else {
				c.Println("Invalid ID format")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows tweets from a user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Insert username: ")
			user := c.ReadLine()
			tweets := tweetManager.GetTweetsByUser(user)
			if tweets == nil {
				c.Println("No tweet with matching user")
			} else {
				for _, tweet := range tweets {
					c.Println(tweet)
				}
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUserTimeline",
		Help: "Shows the timeline of a user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Insert username: ")
			user := c.ReadLine()
			tweets := tweetManager.GetTimeline(user)
			if tweets == nil {
				c.Println("No tweets in timeline")
			} else {
				for _, tweet := range tweets {
					c.Println(tweet)
				}
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts the number of tweets a user has published",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			var user string
			c.Print("Insert user: ")
			user = c.ReadLine()
			count := tweetManager.CountTweetsByUser(user)
			println("The user", user, "has ", count, "tweets")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "followUser",
		Help: "Adds a user to the followed list",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Insert user: ")
			user := c.ReadLine()
			c.Print("Insert user to follow: ")
			followed := c.ReadLine()
			tweetManager.Follow(user, followed)
			println("The user", user, "follows ", followed, "'s tweets")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweets := tweetManager.GetTweets()
			if tweets == nil {
				c.Println("No hay ningun tweet")
			} else {
				for i := 0; i < len(tweets); i++ {
					c.Println(tweets[i])
				}
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweet",
		Help: "Cleans the last tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweetManager.CleanTweet()

			c.Println("Ultimo tweet eliminado")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTrendingTopics",
		Help: "Show the 2 most used words in Tweeter",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			trending := tweetManager.GetTrendingTopics()
			if trending[0] == "" {
				c.Println("No hay trending topics")
			} else {
				for i := 0; i < len(trending); i++ {
					c.Println("#" + trending[i])
				}
			}
		},
	})
	shell.Run()
}
