package main

import (
	"fmt"
	"os"
	"stories/1_echoes_of_enigmatus/quest1"
	"stories/1_echoes_of_enigmatus/quest2"
	"stories/1_echoes_of_enigmatus/quest3"
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
		case 3:
			quest3.Run()
		}
	} else {
		fmt.Println("Quest 1")
		quest1.Run()
		fmt.Println("Quest 2")
		quest2.Run()
		fmt.Println("Quest 3")
		quest3.Run()
	}
}
