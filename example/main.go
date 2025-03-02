package main

import (
	"log"

	"github.com/MisterKirill/go-mediamtx/mediamtx"
)

func main() {
	mediamtx, err := mediamtx.New("http://localhost:9997")
	if err != nil {
		log.Fatal(err)
	}

	paths, err := mediamtx.GetPaths()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d people are streaming right now", len(paths))
}
