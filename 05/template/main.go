package main

import (
	"log"
	"os"
)

func openFile(s string) string {
	constents, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(constents)
}

func main() {
  s := openFile("template")
}
