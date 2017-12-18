package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	var git Versioning = new(Git)

	previous, err := git.CurrentVersion()
	if err != nil {
		log.Fatal(err)
	}

	if !previous.Dirty {
		log.Fatalln("No changes since last version:", previous)
		return
	}

	increment := ""
	if args := flag.Args(); len(args) > 0 {
		increment = args[0]
	}

	current := previous.Increment(increment)

	err = git.UpdateVersion(current)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v -> %v", previous, current)
}
