package quest7

import (
	"fmt"
	"loader"
	"sort"
	"strings"
)

type Chariot struct {
	name      string
	changes   []int
	change_no int
	power     int
	essence   int
}

func parse_data(data []string) []*Chariot {
	chariots := make([]*Chariot, len(data))

	for i, line := range data {
		parts := strings.Split(line, ":")
		name := parts[0]
		change_strs := strings.Split(parts[1], ",")
		changes := make([]int, len(change_strs))
		for j, change := range change_strs {
			switch change {
			case "+":
				changes[j] = 1
			case "-":
				changes[j] = -1
			case "=":
				changes[j] = 0
			}
		}
		chariots[i] = &Chariot{
			name:    name,
			changes: changes,
			power:   10,
		}
	}

	return chariots
}

func parse_grid(grid []string) string {
	height := len(grid)
	width := len(grid[0])

	track := []byte{}
	x, y := 1, 0
	dir := 'R'

	for grid[y][x] != 'S' {
		track = append(track, grid[y][x])

		if dir != 'D' && y > 0 && grid[y-1][x] != ' ' {
			y--
			dir = 'U'
		} else if dir != 'U' && y < height-1 && grid[y+1][x] != ' ' {
			y++
			dir = 'D'
		} else if dir != 'L' && x < width-1 && grid[y][x+1] != ' ' {
			x++
			dir = 'R'
		} else if dir != 'R' && x > 0 && grid[y][x-1] != ' ' {
			x--
			dir = 'L'
		}
	}

	track = append(track, 'S')
	return string(track)
}

func (c *Chariot) move(override byte) {
	if override == '+' {
		c.power++
	} else if override == '-' {
		c.power--
	} else {
		c.power += c.changes[c.change_no]
	}
	if c.power < 0 {
		c.power = 0
	}
	c.essence += c.power

	c.change_no++
	if c.change_no == len(c.changes) {
		c.change_no = 0
	}
}

func race(chariots []*Chariot, turns int) string {
	for i := 0; i < turns; i++ {
		for _, chariot := range chariots {
			chariot.move('=')
		}
	}

	sort.Slice(chariots, func(i, j int) bool {
		return chariots[i].essence > chariots[j].essence
	})

	order := make([]string, len(chariots))
	for i, chariot := range chariots {
		order[i] = chariot.name
	}

	return strings.Join(order, "")
}

func race_track(chariots []*Chariot, track string, loops int) string {
	for _, chariot := range chariots {
		for i := 0; i < loops; i++ {
			for j := 0; j < len(track); j++ {
				chariot.move(track[j])
			}
		}
	}

	sort.Slice(chariots, func(i, j int) bool {
		return chariots[i].essence > chariots[j].essence
	})

	order := make([]string, len(chariots))
	for i, chariot := range chariots {
		order[i] = chariot.name
	}

	return strings.Join(order, "")
}

// Knuth, Donald (2011), "Section 7.2.1.2: Generating All Permutations",
// The Art of Computer Programming, volume 4A.
func next_permutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func try_plan(plan []int, track string, loops int, score_to_beat int) bool {
	chariot := &Chariot{power: 10, changes: plan}

	for i := 0; i < loops; i++ {
		for j := 0; j < len(track); j++ {
			chariot.move(track[j])
		}
	}

	if chariot.essence > score_to_beat {
		return true
	} else {
		return false
	}
}

func try_all_plans(track string, loops int, score_to_beat int) int {
	winning_plans := 0

	plan := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1, 1, 1}
	if try_plan(plan, track, loops, score_to_beat) {
		winning_plans++
	}

	for i := 1; next_permutation(sort.IntSlice(plan)); i++ {
		if try_plan(plan, track, loops, score_to_beat) {
			winning_plans++
		}
	}

	return winning_plans
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 7, 1

	data := loader.GetStrings()
	chariots := parse_data(data)
	part1 := race(chariots, 10)

	loader.Part = 2
	data = loader.GetStrings()
	chariots = parse_data(data)
	grid := []string{
		"S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--",
		"-                                                                     -",
		"=                                                                     =",
		"+                                                                     +",
		"=                                                                     +",
		"+                                                                     =",
		"=                                                                     =",
		"-                                                                     -",
		"--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-",
	}
	track := parse_grid(grid)
	part2 := race_track(chariots, track, 10)

	loader.Part = 3
	data = loader.GetStrings()
	chariots = parse_data(data)
	grid = []string{
		"S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--",
		"- + +   + =   =     =      =   == = - -     - =  =         =-=        -",
		"= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++",
		"+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=       ",
		"= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =          ",
		"+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==",
		"=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =",
		"-               = = = =   +  +  ==+ = = +   =        ++    =          -",
		"-               = + + =   +  -  = + = = +   =        +     =          -",
		"--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-",
	}
	track = parse_grid(grid)
	race_track(chariots, track, 2024)
	rival := chariots[0].essence
	part3 := try_all_plans(track, 2024, rival)

	fmt.Printf("%s %s %d\n", part1, part2, part3)
}
