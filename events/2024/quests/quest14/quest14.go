package quest14

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	x int
	y int
	z int
}

func max_height(data []byte) int {
	height := 0
	max := 0
	steps := strings.Split(string(data), ",")
	for _, step := range steps {
		if step[0] == 'U' || step[0] == 'D' {
			n, err := strconv.Atoi(step[1:])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error %v\n", err)
				os.Exit(1)
			}
			if step[0] == 'U' {
				height += n
				if height > max {
					max = height
				}
			} else if step[0] == 'D' {
				height -= n
			}
		}
	}
	return max
}

func unique_segments(data []string) int {
	segments := map[Segment]struct{}{}

	for _, line := range data {
		segment := Segment{}

		for _, step := range strings.Split(line, ",") {
			dir := step[0]
			n, err := strconv.Atoi(step[1:])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error %v\n", err)
				os.Exit(1)
			}

			for i := 0; i < n; i++ {
				switch dir {
				case 'U':
					segment.y++
				case 'D':
					segment.y--
				case 'R':
					segment.x++
				case 'L':
					segment.x--
				case 'F':
					segment.z++
				case 'B':
					segment.z--
				}
				segments[segment] = struct{}{}
			}
		}
	}

	return len(segments)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 14, 1

	data := loader.GetOneLine()
	part1 := max_height(data)

	loader.Part = 2
	data2 := loader.GetStrings()
	part2 := unique_segments(data2)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
