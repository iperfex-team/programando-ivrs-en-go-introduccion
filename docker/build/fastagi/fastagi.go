package main

import (
	"log"
	"os"

	"github.com/CyCoreSystems/agi"
)


var portFastAGI string
var serverMediaURL string

func init(){

	portFastAGI = goDotEnvVariable("FASTAGI_PORT")
	if portFastAGI == "" {
		portFastAGI = "8000"
	}

	serverMediaURL = goDotEnvVariable("SERVER_MEDIA")
	if serverMediaURL == "" {
		serverMediaURL = "http://fastagi:8011/"
	}

}

func main() {

	log.Println("FastAGI Server agi://0.0.0.0:" + portFastAGI + " running...")
	agi.Listen(":"+portFastAGI, handler)

}

func handler(a *agi.AGI) {
	defer a.Close()

	a.Verbose(". .: DEMO :.", 1)
	a.Answer()
	a.Set("CHANNEL(language)", "es")
	a.Verbose("PRUEBA 4 - PLAYBACK AUDIO", 1)
	a.StreamFile("tt-monkeys", "#", 1)
	a.Set("__VAR_GLOBAL", a.Variables["agi_arg_1"])
	a.Hangup()

}