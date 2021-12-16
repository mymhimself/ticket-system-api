package main

import (
	"github.com/mymhimself/ticket-system-api/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
