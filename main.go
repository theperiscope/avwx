package main

import (
	"log"
	"os"

	"github.com/theperiscope/avwx/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
