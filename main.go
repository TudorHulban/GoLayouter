package main

import (
	"log"
	"os"

	"github.com/TudorHulban/GoLayouter/helpers"
	"github.com/TudorHulban/GoLayouter/interfaces"
	"github.com/TudorHulban/GoLayouter/objects"
)

func main() {
	fileSource := os.Args[1]

	content, err := helpers.ReadFile(fileSource)
	if err != nil {
		log.Print(err)
	}

	e := objects.NewEntries(content)
	entries := e.Parse()

	writter := objects.ConvertToIWritter(entries)
	err = interfaces.Write(writter)
	if err != nil {
		log.Print(err)
	}
}
