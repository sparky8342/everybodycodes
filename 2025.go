package main

import (
	"2025/quest1"
	"2025/quest2"
	"2025/quest3"
	"2025/quest4"
	"fmt"
	"os"
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
		case 4:
			quest4.Run()
		}
	} else {
		fmt.Println("Quest 1")
		quest1.Run()
		fmt.Println("Quest 2")
		quest2.Run()
		fmt.Println("Quest 3")
		quest3.Run()
		fmt.Println("Quest 4")
		quest4.Run()
	}
}
