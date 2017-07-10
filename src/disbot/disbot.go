package disbot

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"log"
	"github.com/fatih/color"
)

type JsonObject struct {

	Token string
	TokenKey string
	BotId int

}


func Connect(file []byte){

	var config JsonObject

	json.Unmarshal(file,&config)

	if config.TokenKey == "" {
		fmt.Println("Start Discord Bot")
		return
	}



	dBot , err := discordgo.New("Bot " + config.TokenKey)
	if err != nil {
		panic(err)
	}
	err = dBot.Open()
	if err != nil {
		panic(err)
	}

	// EventHandler

	dBot.AddHandler(ready)
	dBot.AddHandler(createMessage)
	dBot.AddHandler(UpdateGuild)



// Start Discord Bot
	color.Set(color.FgGreen)
	log.SetPrefix("PandaBot: ")
	log.Println("Panda Bot is now running.  Press CTRL-C to exit.")
	color.Set(color.FgWhite)


	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dBot.Close()

}

