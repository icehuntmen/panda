package main
import (

	"log"
	"github.com/fatih/color"

)



func logTrace(v...interface{})  {
	color.Set(color.FgYellow)
	log.SetPrefix("TRACE: ")
	log.Println(v...)
	color.Set(color.FgWhite)
}

func logInfo(v...interface{})  {
	color.Set(color.FgGreen)
	log.SetPrefix("INFO: ")
	log.Println(v...)
	color.Set(color.FgWhite)
}

func logSystem(v...interface{})  {
	color.Set(color.FgGreen)
	log.SetPrefix("PandaBot: ")
	log.Println(v...)
	color.Set(color.FgWhite)
}

func logError(v...interface{})  {
	color.Set(color.FgRed)
	log.SetPrefix("ERROR: ")
	log.Println(v...)
	color.Set(color.FgWhite)
}

func panicOnErr(err error) {
	if err != nil {
		color.Set(color.FgHiRed)
		panic(err)
		color.Set(color.FgWhite)
	}
}