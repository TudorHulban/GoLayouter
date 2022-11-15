package main

import (
	"log"
	"os"

	"github.com/TudorHulban/GoLayouter/domain/interfaces"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

func main() {
	fileSource := os.Args[1]

	content, err := interfaces.IRWritterReadFile(fileSource)
	if err != nil {
		log.Print(err)
	}

	e := objects.NewEntries(content)
	entries := e.Parse()

	writter := interfaces.ConvertToIWritter(entries)
	err = interfaces.WriteToDisk(writter)
	if err != nil {
		log.Print(err)
	}
}
