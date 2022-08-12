package main

import (
	"log"
	"musement/src/internal/app"
)

func recoverPanicOnStartup() {
	if r := recover(); r != nil {
		log.Panicf("panic on startup ='%v'", r)
	}
}

func main() {
	defer recoverPanicOnStartup()

	err := app.StartApp()

	if err != nil {
		log.Println("error starting app")
	}
}
