package main

import (
	"events/2024/quests/quest1"
	"events/2024/quests/quest2"
	"events/2024/quests/quest3"
	"events/2024/quests/quest4"
	"fmt"
)

func main() {
	fmt.Println("Quest 1")
	quest1.Run()
	fmt.Println("Quest 2")
	quest2.Run()
	fmt.Println("Quest 3")
	quest3.Run()
	fmt.Println("Quest 4")
	quest4.Run()
}
