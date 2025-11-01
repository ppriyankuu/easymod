package main

import (
	"easymod/cmd"
	"log"
	"os"
)

func main() {
	app := cmd.Execute()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
