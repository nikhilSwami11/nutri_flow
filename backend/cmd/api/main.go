package main

import (
	"log"

	"github.com/nikhilswami11/nutriflow/backend/cmd/api/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
