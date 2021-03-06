package main

import (
	"log"
	"os"

	"github.com/chck/synr/chatwork"
	"github.com/chck/synr/slack"

	"github.com/chck/synr/config"
	flags "github.com/jessevdk/go-flags"
)

type options struct {
	ChatName    string `short:"c" long:"chatname" description:"A chat name you'd like to leave"`
	DryRun      bool   `short:"d" long:"dry-run" description:"Pre-running to leave unnessesary chat rooms"`
	BeforeMonth int    `short:"m" long:"before-month" description:"Set X month elapsed when Last of talking date to leave: DEFAULT 1 MONTH AGO"`
}

func cmdOpts() *options {
	opts := &options{}
	parser := flags.NewParser(opts, flags.PrintErrors)
	parser.Name = "synr"
	parser.Usage = "-c slack"
	args, _ := parser.Parse()

	if len(args) != 0 || opts.ChatName == "" {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	return opts
}

func main() {
	log.Println("Erase your past glory...")
	log.Println("++++++++++++++++++++++++")
	tokens := config.Load().Tokens

	opts := cmdOpts()

	switch opts.ChatName {
	case "chatwork":
		client := chatwork.New(tokens.Chatwork)
		rooms, _ := client.GetRooms()
		for _, room := range rooms {
			chatwork.MayBeLeaveRoom(opts.DryRun, opts.BeforeMonth, client, &room)
		}
	case "slack":
		client := slack.New(tokens.Slack)
		channels, _ := client.GetChannels(false)
		starredIDs := slack.StarredChannelIDs(client)
		for _, channel := range channels {
			slack.MayBeLeaveChannel(opts.DryRun, opts.BeforeMonth, client, channel, starredIDs)
		}
	}
	log.Println("++++++++++++++++++++++++")
	log.Println("Done!!")
}
