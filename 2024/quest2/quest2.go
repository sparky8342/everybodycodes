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

	count := 0
	for _, symbol := range symbols {
		if symbol {
			count++
		}
	}
	return count
}

func add_wrap(n int, inc int, limit int) int {
	n += inc
	if n == limit {
		return 0
	} else if n == -1 {
		return limit - 1
	} else {
		return n
	}
}

func scales(words []string, phrase string) int {
	scales := make([]bool, len(phrase))

	grid := strings.Split(phrase, "\n")
	height := len(grid)
	width := len(grid[0])

	dirs := [][]int{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
	}

	for _, word := range words {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
			dir_loop:
				for _, dir := range dirs {
					if grid[y][x] == word[0] {

						search_x := x
						search_y := y
						for i := 1; i < len(word); i++ {
							search_y += dir[1]
							if search_y == -1 || search_y == height {
								continue dir_loop
							}
							search_x = add_wrap(search_x, dir[0], width)

							if grid[search_y][search_x] != word[i] {
								continue dir_loop
							}
						}

						search_x = x
						search_y = y
						for i := 0; i < len(word); i++ {
							scales[search_y*width+search_x] = true
							search_y += dir[1]
							search_x = add_wrap(search_x, dir[0], width)
						}
					}
				}
			}
		}
	}

	count := 0
	for _, scale := range scales {
		if scale {
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

	loader.Part = 3
	data = loader.GetStrings()
	words, phrase = parse_data(data)
	part3 := scales(words, phrase)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
