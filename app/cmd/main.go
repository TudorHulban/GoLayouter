package main

import (
	"log"
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/app/service"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

func main() {
	fileSource := os.Args[1]

	content, err := helpers.ReadFile(fileSource)
	if err != nil {
		log.Print(err)
	}

	entries := objects.NewEntries(content).Parse()
	serv := service.NewService(entries)

	err = serv.WriteToDisk()
	if err != nil {
		log.Print(err)
	}
}
