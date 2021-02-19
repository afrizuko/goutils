package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hello World!")
	l := log.New(os.Stderr, "", 0)
	l.Println("Hello World from l")
}
