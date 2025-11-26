package quest17

import (
	"container/heap"
	"fmt"
	"loader"
	"math"
)

type Pos struct {
	x int
	y int
}

type Entry struct {
	pos         Pos
	visited_bot int
	distance    int
}

var dirs [4][2]int

func init() {
	dirs = [4][2]int{
		[2]int{1, 0},
		[2]int{-1, 0},
		[2]int{0, 1},
		[2]int{0, -1},
	}
}

type PriorityQueue []*Entry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Entry)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func pow2(n int) int {
	return n * n
}

func lava_spread(grid []string, radius int) int {
	size := len(grid)

	cv := size / 2
	square_radius := radius * radius

	total := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if x == cv && y == cv {
				fmt.Print("@")
			} else if pow2(cv-x)+pow2(cv-y) <= square_radius {
				total += int(grid[y][x] - '0')
				fmt.Print(string(grid[y][x]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return total
}

func max_ring(grid []string) int {
	size := len(grid)

	cv := size / 2

	max := 0
	max_radius := 0

	for radius := 1; radius <= cv; radius++ {
		square_radius := radius * radius
		prev_square_radius := pow2(radius - 1)

		lava := 0
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				if x == cv && y == cv {
					continue
				}

				r := pow2(cv-x) + pow2(cv-y)
				if r <= square_radius && r > prev_square_radius {
					lava += int(grid[y][x] - '0')
				}
			}
		}

		if lava > max {
			max = lava
			max_radius = radius
		}
	}

	return max * max_radius
}

func find_start(grid [][]byte) Pos {
	size := len(grid)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x] == 'S' {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{}
}

func best_loop(data []string) int {
	size := len(data)
	cv := size / 2

	grid := make([][]byte, size)
	for i := range grid {
		grid[i] = []byte(data[i])
	}

	start := find_start(grid)

	for radius := 1; ; radius++ {
		distance := shortest_path(grid, start)
		if distance < radius*30 {
			return distance * (radius - 1)
		}

		// clear spaces hit by lava
		square_radius := radius * radius
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				if x == cv && y == cv {
					continue
				}
				if pow2(cv-x)+pow2(cv-y) <= square_radius {
					grid[y][x] = '.'
				}
			}
		}
	}

	return -1
}

func shortest_path(grid [][]byte, start Pos) int {
	size := len(grid)
	cv := size / 2

	distances := make([][][]int, size)
	for y := range distances {
		row := make([][]int, size)
		for x := 0; x < size; x++ {
			row[x] = []int{math.MaxInt32, 0}
		}
		distances[y] = row
	}

	distances[start.y][start.x] = []int{0, 0}

	queue := make(PriorityQueue, 1)
	queue[0] = &Entry{pos: start}
	heap.Init(&queue)

	for queue.Len() > 0 {
		entry := heap.Pop(&queue).(*Entry)

		pos := entry.pos

		if entry.visited_bot == 1 && pos.x == start.x && pos.y == start.y {
			return entry.distance
		}

		for _, dir := range dirs {
			new_pos := Pos{x: pos.x + dir[0], y: pos.y + dir[1]}
			if new_pos.x < 0 || new_pos.x == size || new_pos.y < 0 || new_pos.y == size {
				continue
			}
			if grid[new_pos.y][new_pos.x] != '.' && grid[new_pos.y][new_pos.x] != '@' {
				new_visited_bot := entry.visited_bot
				if entry.visited_bot == 0 && new_pos.x > cv {
					continue
				}
				if entry.visited_bot == 0 && new_pos.x == cv && new_pos.y > cv {
					new_visited_bot = 1
				}

				dist := entry.distance
				if grid[new_pos.y][new_pos.x] != 'S' {
					dist += int(grid[new_pos.y][new_pos.x] - '0')
				}

				if dist < distances[new_pos.y][new_pos.x][0] || (new_visited_bot > distances[new_pos.y][new_pos.x][1] && new_pos.y < cv) {
					distances[new_pos.y][new_pos.x] = []int{dist, new_visited_bot}
					heap.Push(&queue, &Entry{pos: new_pos, distance: dist, visited_bot: new_visited_bot})
				}
			}
		}
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 17, 1

	grid := loader.GetStrings()
	part1 := lava_spread(grid, 10)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := max_ring(grid)

	loader.Part = 3
	grid = loader.GetStrings()
	part3 := best_loop(grid)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
