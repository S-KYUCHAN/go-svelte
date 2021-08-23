package main

import (
	// "fmt"
	// "log"
	// "os"
	// "os/signal"
	
	"github.com/S-KYUCHAN/backend/engine"
)


func main() {

	engine.HandleRequests()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Print("closing database connection")
}