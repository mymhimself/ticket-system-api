package main

import (
	"fmt"
	"github.com/mymhimself/ticket-system-api/cmd"
	"log"
)

func main() {
	fmt.Println("salaam Mr Ehsani")

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
