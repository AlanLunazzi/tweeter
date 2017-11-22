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
			service.PublishTweet(tweet)
			c.Print("Tweet Sent\n")
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
