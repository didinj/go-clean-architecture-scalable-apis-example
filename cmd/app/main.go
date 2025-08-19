package main

import (
	"log"

	"github.com/didinj/go-clean-architecture/internal/bootstrap"
)

func main() {
	app := bootstrap.InitializeApp()
	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
