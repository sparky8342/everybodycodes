package quest3

import (
	"fmt"
	"loader"
	"regexp"
	"strconv"
	"strings"
)

type Die struct {
	faces     []int
	seed      int
	pulse     int
	roll_no   int
	current   int
	track_pos int
}

type Pos struct {
	x int
	y int
}

var dirs [][]int
var width, height int

var cache map[[3]int]struct{}

func init() {
	dirs = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{0, 0},
	}
}

func parse_data(data []string, mode int) ([]Die, []int, [][]int) {
	r := regexp.MustCompile(`faces=\[(.*?)\] seed=(\d+)`)

	track := []int{}
	grid := [][]int{}

	if mode == 2 {
		track_str := data[len(data)-1]
		data = data[:len(data)-2]
		for _, ru := range track_str {
			track = append(track, int(ru-'0'))
		}
	} else if mode == 3 {
		var grid_data []string
		for i, line := range data {
			if line == "" {
				grid_data = data[i+1:]
				data = data[0:i]
				break
			}
		}
		for _, line := range grid_data {
			row := []int{}
			for _, ru := range line {
				row = append(row, int(ru-'0'))
			}
			grid = append(grid, row)
		}
	}

	dice := make([]Die, len(data))

	for i, line := range data {
		die := Die{}

		matches := r.FindStringSubmatch(line)

		for _, str := range strings.Split(matches[1], ",") {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			die.faces = append(die.faces, n)
		}

		seed, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		die.seed = seed
		die.pulse = seed
		die.roll_no = 1

		dice[i] = die
	}

	return dice, track, grid
}

func (d *Die) roll() int {
	spin := d.roll_no * d.pulse
	d.pulse += spin
	d.pulse %= d.seed
	d.pulse += 1 + d.roll_no + d.seed
	d.current = (d.current + spin) % len(d.faces)
	d.roll_no++
	return d.faces[d.current]
}

func (d *Die) clone() *Die {
	new_d := Die{
		faces:     d.faces,
		seed:      d.seed,
		pulse:     d.pulse,
		roll_no:   d.roll_no,
		current:   d.current,
		track_pos: d.track_pos,
	}
	return &new_d
}

func roll_dice(dice []Die, target int) int {
	score := 0
	rolls := 0
	for score < target {
		for i := range dice {
			score += dice[i].roll()
		}
		rolls++
	}
	return rolls
}

func race_track(dice []Die, track []int) string {
	order := []string{}
	for len(order) < len(dice) {
		for i := range dice {
			if dice[i].track_pos < len(track) {
				roll := dice[i].roll()
				if roll == track[dice[i].track_pos] {
					dice[i].track_pos++
					if dice[i].track_pos == len(track) {
						order = append(order, strconv.Itoa(i+1))
					}
				}
			}
		}
	}
	return strings.Join(order, ",")
}

func search(die *Die, x int, y int, grid [][]int, tokens *[][]bool) {
	key := [3]int{die.roll_no, x, y}
	if _, ok := cache[key]; ok {
		return
	}
	cache[key] = struct{}{}

	roll := die.roll()

	new_positions := []Pos{}
	for _, dir := range dirs {
		new_x := x + dir[0]
		new_y := y + dir[1]
		if new_x < 0 || new_x == width || new_y < 0 || new_y == height {
			continue
		}
		if grid[new_y][new_x] == roll {
			(*tokens)[new_y][new_x] = true
			new_positions = append(new_positions, Pos{x: new_x, y: new_y})
		}
	}

	if len(new_positions) == 1 {
		pos := new_positions[0]
		search(die, pos.x, pos.y, grid, tokens)
	} else {
		for _, pos := range new_positions {
			search(die.clone(), pos.x, pos.y, grid, tokens)
		}
	}
}

func grid_game(dice []Die, grid [][]int) int {
	height = len(grid)
	width = len(grid[0])

	tokens := make([][]bool, height)
	for i := range tokens {
		tokens[i] = make([]bool, width)
	}

	for i := range dice {
		cache = map[[3]int]struct{}{}

		roll := dice[i].roll()
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == roll {
					tokens[y][x] = true
					search(dice[i].clone(), x, y, grid, &tokens)
				}
			}
		}
	}

	total := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if tokens[y][x] {
				total++
			}
		}
	}

	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2", 3, 1

	data := loader.GetStrings()
	dice, _, _ := parse_data(data, 1)
	part1 := roll_dice(dice, 10000)

	loader.Part = 2
	data = loader.GetStrings()
	dice, track, _ := parse_data(data, 2)
	part2 := race_track(dice, track)

	loader.Part = 3
	data = loader.GetStrings()
	dice, _, grid := parse_data(data, 3)
	part3 := grid_game(dice, grid)

	fmt.Printf("%d\n%s\n%d\n", part1, part2, part3)
}
