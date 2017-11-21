package service

var tweet string = ""

func PublishTweet(tw string) {
	tweet = tw

}

func GetTweet() string {
	return tweet
}
