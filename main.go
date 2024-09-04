package main

import (
	"log"

	"github.com/ezequieljn/morse-code/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
