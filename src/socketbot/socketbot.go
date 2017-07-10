package socketbot

import (
	"log"
	"github.com/fatih/color"
)

func WebConnection()  {

	color.Set(color.FgGreen)
	log.SetPrefix("WebSocket: ")
	log.Println("Web is now running.")
	color.Set(color.FgWhite)
}
