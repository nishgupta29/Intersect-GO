package main

#Importing required modules

import (
	"os"

	"github.com/shomali11/slacker"
)
#Request writer

func handle(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("Hey!")
}
#Connecting to slack and providing a response

func main() {
	bot := slacker.NewClient(os.Getenv("API_TOKEN"))
	bot.Command("hello", "Say hello", handle)
	err := bot.Listen()
	if err != nil {
		panic(err)
	}
}
