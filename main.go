package main 

import (
	"fmt"
	"log"
	"context"
	"os"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6177862970372-6198505808912-lkopW0bwLCMInXoPksoxKSP4")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0657U125B6-6175406421523-89c0894f26985bf451479ffa2d538adee03a9f45da05edd1a32e86c501f538fa")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	
}