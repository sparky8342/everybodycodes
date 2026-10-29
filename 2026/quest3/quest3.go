package quest3

import (
	"fmt"
	"loader"
	"strings"
)

func pow(n, p int) int {
	ans := n
	for i := 1; i < p; i++ {
		ans *= n
	}
	return ans
}

func score(bases []int, number []int) int {
	sc := number[len(number)-1]
	p := 1
	var base_pos int
	if len(bases) == 1 {
		base_pos = 0
	} else {
		base_pos = len(bases) - 2
	}
	for i := len(number) - 2; i >= 0; i-- {
		sc += pow(bases[base_pos], p) * number[i]
		p++
		base_pos--
		if base_pos == -1 {
			base_pos = len(bases) - 1
		}
	}
	return sc
}

func add1(num []int) []int {
	p := len(num) - 1
	num[p]++
	for num[p] == 10 {
		num[p] = 0
		p--
		if p == -1 {
			num = append([]int{0}, num...)
			p = 0
		}
		num[p]++
	}
	return num
}

func comp(num1 []int, num2 []int) bool {
	if len(num1) != len(num2) {
		return false
	}
	for i := 0; i < len(num1); i++ {
		if num1[i] != num2[i] {
			return false
		}
	}
	return true
}

func find_smaller(bases []int, number []int, sc int) bool {
	num := []int{0}
	for {
		current_score := score(bases, num)
		if current_score == sc {
			return true
		}
		num = add1(num)
		if comp(num, number) {
			return false
		}
	}
}

func total_score(data []string) int {
	total := 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		sub_parts := strings.Split(parts[0], "-")
		iso := sub_parts[0]
		bases := []int{}
		for _, d := range sub_parts[1] {
			bases = append(bases, int(d-'0'))
		}
		number := []int{}
		for _, d := range parts[1] {
			number = append(number, int(d-'0'))
		}

		if iso == "ISO2" {
			sc := score(bases, number)
			if !find_smaller(bases, number, sc) {
				total += sc
			}
		} else {
			total += score(bases, number)
		}
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 3, 1

	data := loader.GetStrings()
	part1 := total_score(data)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := total_score(data)

	loader.Part = 3
	data = loader.GetStrings()
	part3 := total_score(data)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
