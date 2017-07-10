package disbot

import (
	"github.com/bwmarrin/discordgo"
	"strings"

	"log"
)

func ready(s *discordgo.Session,  event *discordgo.Ready){

	s.UpdateStatus(0, "Overwatch")
}



func createMessage(s *discordgo.Session, event *discordgo.MessageCreate){


	if event.Author.ID == s.State.User.ID || event.Author.Bot{
		return
	}

	cLower := strings.ToLower(event.Content)
	cVoce := event.Content



	log.Println(s.State.SessionID)
	log.Println(event.Message)
	log.Printf("ID: %v",event.ID)
	log.Printf("Channel ID: %v",event.ChannelID)
	log.Printf("Author: %v",event.Author)
	log.Printf("Content: %v",event.Content)
	log.Printf("Timestamp: %v",event.Timestamp)
	log.Printf("EditedTimestamp: %v",event.EditedTimestamp)
	log.Printf("MentionRoles: %v",event.MentionRoles)
	log.Printf("Tts: %v",event.Tts)
	log.Printf("MentionEveryone: %v",event.MentionEveryone)
	log.Printf("Embeds: %v",event.Embeds)
	log.Printf("Mentions: %v",event.Mentions)
	log.Printf("Reactions: %v",event.Reactions)


	if strings.HasPrefix(cLower, "who's a good boy") {
		s.ChannelMessageSend(event.ChannelID, "ME ME ME <@"+event.Author.ID+">")
	}

	if strings.HasPrefix(cLower, "test") {
		s.ChannelMessageSend(event.ChannelID, cVoce)
	}

}

func UpdateGuild(s *discordgo.Session, event *discordgo.PresenceUpdate){

	log.Println(s,event)
}

func PrivateMessage(v...interface{})  {

}