package main

import (
	"fmt"
	"net/http"

	TwitchApi "github.com/moosetheory/ocebot/Twitch"
)

var (
	clientID     = "p1sphibop7dvhv4kienpk6krfzh98u"
	clientSecret = "9113iare0j59axn9jq3qu81rspa57e"
	accessToken  string
)

func main() {
	http.HandleFunc("/stream/online", handleStreamStartNotifiction)
	twitchConneciton := TwitchApi.TwitchApi{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	err := twitchConneciton.Connect()
	if err != nil {
		panic(err)
	}
	twitchConneciton.GetExistingSubs()
	twitchConneciton.PrintInfo()
}

func handleStreamStartNotifiction(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
