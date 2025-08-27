package quest1

import (
	"fmt"
	"loader"
	"sort"
)

type Play struct {
	slot  int
	score int
}

var height, width int

func parse_data(data []string) ([]string, []string) {
	var grid, tokens []string

	for i, row := range data {
		if row == "" {
			grid = data[0:i]
			tokens = data[i+1:]
			break
		}
	}

	height = len(grid)
	width = len(grid[0])

	return grid, tokens
}

func play_token(grid []string, slot int, token string) int {
	x := (slot - 1) * 2
	y := -1

	token_i := 0

	for {
		for grid[y+1][x] == '.' {
			y++
			if y == height-1 {
				final_slot := x/2 + 1
				score := final_slot*2 - slot
				if score < 0 {
					score = 0
				}
				return score
			}
		}

		if x == 0 {
			x++
		} else if x == width-1 {
			x--
		} else if token[token_i] == 'R' {
			x++
		} else if token[token_i] == 'L' {
			x--
		}

		token_i++
	}

	return 0
}

func play_tokens(grid []string, tokens []string) int {
	total_score := 0

	for i := 0; i < len(tokens); i++ {
		slot := i + 1
		total_score += play_token(grid, slot, tokens[i])
	}

	return total_score
}

func maximise_tokens(grid []string, tokens []string) int {
	slots := (width + 1) / 2
	total_score := 0

	for _, token := range tokens {
		max_score := 0
		for i := 1; i <= slots; i++ {
			score := play_token(grid, i, token)
			if score > max_score {
				max_score = score
			}
		}
		total_score += max_score
	}

	return total_score
}

func search(scores [][]Play, pick []int, best_score *int, dir int) {
	clashes := map[int]struct{}{}

	slot_numbers := map[int]int{}
	for i := 0; i < len(scores); i++ {
		slot := scores[i][pick[i]].slot
		if val, ok := slot_numbers[slot]; !ok {
			slot_numbers[slot] = i
		} else {
			clashes[i] = struct{}{}
			clashes[val] = struct{}{}
		}
	}

	if len(clashes) == 0 {
		score := 0
		for i := 0; i < len(scores); i++ {
			score += scores[i][pick[i]].score
		}
		if dir == 1 && score > *best_score {
			*best_score = score
		} else if dir == -1 && score < *best_score {
			*best_score = score
		}
	} else {
		for clash := range clashes {
			pick[clash] += dir
			search(scores, pick, best_score, dir)
			pick[clash] -= dir
		}
	}
}

func unique_slots(grid []string, tokens []string) string {
	slots := (width + 1) / 2

	scores := make([][]Play, len(tokens))

	for i, token := range tokens {
		token_scores := make([]Play, slots)
		for j := 0; j < slots; j++ {
			score := play_token(grid, j+1, token)
			token_scores[j] = Play{slot: j + 1, score: score}
		}
		sort.Slice(token_scores, func(i, j int) bool {
			return token_scores[i].score > token_scores[j].score
		})
		scores[i] = token_scores
	}

	pick := make([]int, len(tokens))
	max_score := 0
	search(scores, pick, &max_score, 1)

	for i := 0; i < len(tokens); i++ {
		pick[i] = slots - 1
	}
	min_score := 1000
	search(scores, pick, &min_score, -1)

	return fmt.Sprintf("%d %d", min_score, max_score)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2", 1, 1

	data := loader.GetStrings()
	grid, tokens := parse_data(data)
	part1 := play_tokens(grid, tokens)

	loader.Part = 2
	data = loader.GetStrings()
	grid, tokens = parse_data(data)
	part2 := maximise_tokens(grid, tokens)

	loader.Part = 3
	data = loader.GetStrings()
	grid, tokens = parse_data(data)
	part3 := unique_slots(grid, tokens)

	fmt.Printf("%d\n%d\n%s\n", part1, part2, part3)
}
