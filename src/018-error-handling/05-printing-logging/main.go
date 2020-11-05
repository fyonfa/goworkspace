package main

import (
	//"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happeneeeeed:", err) //normal error customed
		log.Println("err happened: ", err) //this one logs it with time

	}
	/*
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bs))

	*/
}
