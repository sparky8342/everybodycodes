package main

import (
	"fmt"
	"os"
	"stories/2_the_entertainment_hub/quest1"
	"stories/2_the_entertainment_hub/quest2"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		part, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		switch part {
		case 1:
			quest1.Run()
		case 2:
			quest2.Run()
		/*
		case 3:
			quest3.Run()
		*/
		}
	} else {
		fmt.Println("Quest 1")
		quest1.Run()
		fmt.Println("Quest 2")
		quest2.Run()
		/*
		fmt.Println("Quest 3")
		quest3.Run()
		*/
	}
}
