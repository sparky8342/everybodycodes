package main

import (
	"2024/quest1"
	"2024/quest10"
	"2024/quest11"
	"2024/quest12"
	"2024/quest13"
	"2024/quest14"
	"2024/quest15"
	"2024/quest16"
	"2024/quest17"
	"2024/quest18"
	"2024/quest19"
	"2024/quest2"
	"2024/quest20"
	"2024/quest3"
	"2024/quest4"
	"2024/quest5"
	"2024/quest6"
	"2024/quest7"
	"2024/quest8"
	"2024/quest9"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		part, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v\n", err)
			os.Exit(1)
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
		case 5:
			quest5.Run()
		case 6:
			quest6.Run()
		case 7:
			quest7.Run()
		case 8:
			quest8.Run()
		case 9:
			quest9.Run()
		case 10:
			quest10.Run()
		case 11:
			quest11.Run()
		case 12:
			quest12.Run()
		case 13:
			quest13.Run()
		case 14:
			quest14.Run()
		case 15:
			quest15.Run()
		case 16:
			quest16.Run()
		case 17:
			quest17.Run()
		case 18:
			quest18.Run()
		case 19:
			quest19.Run()
		case 20:
			quest20.Run()
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
		fmt.Println("Quest 5")
		quest5.Run()
		fmt.Println("Quest 6")
		quest6.Run()
		fmt.Println("Quest 7")
		quest7.Run()
		fmt.Println("Quest 8")
		quest8.Run()
		fmt.Println("Quest 9")
		quest9.Run()
		fmt.Println("Quest 10")
		quest10.Run()
		fmt.Println("Quest 11")
		quest11.Run()
		fmt.Println("Quest 12")
		quest12.Run()
		fmt.Println("Quest 13")
		quest13.Run()
		fmt.Println("Quest 14")
		quest14.Run()
		fmt.Println("Quest 15")
		quest15.Run()
		fmt.Println("Quest 16")
		quest16.Run()
		fmt.Println("Quest 17")
		quest17.Run()
		fmt.Println("Quest 18")
		quest18.Run()
		fmt.Println("Quest 19")
		quest19.Run()
		fmt.Println("Quest 20")
		quest20.Run()
	}
}
