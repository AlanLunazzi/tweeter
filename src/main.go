package main

import (
	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands \n")
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			usr := c.ReadLine()
			c.Print("Write your tweet: ")
			txt := c.ReadLine()
			tweet := domain.NewTweet(usr, txt)
			err := service.PublishTweet(tweet)
			if err != nil {
				c.Println("Error->", err)
			} else {
				c.Println("Tweet Sent")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := service.GetTweet()
			if tweet == nil {
				c.Println("No hay ultimo tweet")
			} else {
				c.Println(tweet.User)
				c.Println(tweet.Text)
				c.Println(tweet.Date)
			}

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweets := service.GetTweets()
			if tweets == nil {
				c.Println("No hay ningun tweet")
			} else {
				for i := 0; i < len(tweets); i++ {
					c.Println("Tweet ", i)
					c.Println(tweets[i].User)
					c.Println(tweets[i].Text)
					c.Println(tweets[i].Date)
				}
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweet",
		Help: "Cleans the last tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			service.CleanTweet()

			c.Println("Ultimo tweet eliminado")
		},
	})

	shell.Run()
}
