package main

import (
	"fmt"
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/app/service"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

// TODO merge branch
// golang cli
// yaml config
// github actions
//

func main() {
	fileSource := os.Args[1]

	content, errRead := helpers.ReadFile(fileSource)
	if errRead != nil {
		fmt.Print(errRead)
		os.Exit(1)
	}

	entries := objects.NewEntries(content).Parse()
	serv, errNewService := service.NewService(entries)
	if errNewService != nil {
		fmt.Print(errNewService)
		os.Exit(2)
	}

	errWrite := serv.WriteToDisk()
	if errWrite != nil {
		fmt.Print(errWrite)
		os.Exit(3)
	}
}
