package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("Lorem.txt")
	if err != nil {
		log.Fatalln("cannot open file", err)
	}
	defer file.Close()
}
