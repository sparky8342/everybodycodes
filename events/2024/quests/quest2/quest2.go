package quest2

import (
	"fmt"
	"loader"
	"strings"
)

func parse_data(data []string) ([]string, string) {
	words := strings.Split(data[0][6:], ",")
	phrase := strings.Join(data[2:], "\n")
	return words, phrase
}

func contains(words []string, phrase string) int {
	count := 0
	for _, word := range words {
		count += strings.Count(phrase, word)
	}
	return count
}

func reverse_string(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse_words(words []string) []string {
	reversed := make([]string, len(words))
	for i, word := range words {
		reversed[i] = reverse_string(word)
	}
	return reversed
}

func runic_symbols(words []string, phrase string) int {
	words = append(words, reverse_words(words)...)

	symbols := make([]bool, len(phrase))

	for _, word := range words {
	outer:
		for i := 0; i < len(phrase); i++ {
			if i+len(word) > len(phrase) {
				break
			}
			for j := 0; j < len(word); j++ {
				if phrase[i+j] != word[j] {
					continue outer
				}
			}
			for j := 0; j < len(word); j++ {
				symbols[i+j] = true
			}
		}
	}

	/*
		fmt.Println(phrase)
		for i := 0; i < len(phrase); i++ {
			if symbols[i] {
				fmt.Print(string(phrase[i]))
			} else {
				fmt.Print(strings.ToLower(string(phrase[i])))
			}
		}
		fmt.Println()
	*/

	count := 0
	for _, symbol := range symbols {
		if symbol {
			count++
		}
	}
	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 2, 1

	data := loader.GetStrings()
	words, phrase := parse_data(data)
	part1 := contains(words, phrase)

	loader.Part = 2
	data = loader.GetStrings()
	words, phrase = parse_data(data)
	part2 := runic_symbols(words, phrase)

	fmt.Printf("%d %d\n", part1, part2)
}
