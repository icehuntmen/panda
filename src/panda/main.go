package main

import (
	"io/ioutil"
	"disbot"
	"socketbot"
)



func main()  {


	// Read json config file
	file, e := ioutil.ReadFile("./config.json")
	panicOnErr(e)

	socketbot.WebConnection()
	disbot.Connect(file)

}
