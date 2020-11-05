package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("err happeneeeeed:", err) //normal error customed
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Err happened here", err)
	}
	defer f2.Close()
	log.SetOutput(f)

}
