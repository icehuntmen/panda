package main

import (

	"io/ioutil"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type JsonObject struct {

	Token string
	TokenKey string
	BotId int

}

var config JsonObject

func main()  {

	// Read json config file
	file, e := ioutil.ReadFile("./config.json")
	panicOnErr(e)
	//
	json.Unmarshal(file,&config)
	if config.TokenKey == "" {
		logError("No token provided. Please set token in to config.json")
		return
	}

	discord, err := discordgo.New(config.TokenKey)
	panicOnErr(err)

	err = discord.Open()
	panicOnErr(err)

	logSystem("Panda Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
