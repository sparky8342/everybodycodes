package quest2

import (
	"fmt"
	"loader"
	"strings"
)

func parse_data(data []string) ([]string, string) {
	words := strings.Split(data[0][6:], ",")
	phrase := data[2]
	return words, phrase
}

func contains(words []string, phrase string) int {
	count := 0
	for _, word := range words {
		count += strings.Count(phrase, word)
	}
	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 2, 1

	data := loader.GetStrings()
	words, phrase := parse_data(data)
	part1 := contains(words, phrase)

	fmt.Println(part1)
}
