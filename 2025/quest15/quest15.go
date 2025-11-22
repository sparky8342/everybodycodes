package quest15

import (
	"container/heap"
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Move struct {
	dir    byte
	amount int
}

type Pos struct {
	x int
	y int
}

type Entry struct {
	pos      Pos
	distance int
}

type Wall struct {
	start    Pos
	end      Pos
	vertical bool
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

var cardinal_dirs [4][2]int

func init() {
	cardinal_dirs = [4][2]int{
		[2]int{1, 0},
		[2]int{-1, 0},
		[2]int{0, 1},
		[2]int{0, -1},
	}
}

func parse_data(data []byte) []Move {
	parts := strings.Split(string(data), ",")
	moves := make([]Move, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part[1:])
		if err != nil {
			panic(err)
		}
		moves[i] = Move{
			dir:    part[0],
			amount: n,
		}
	}
	return moves
}

func find_exit(moves []Move) int {
	dirs := [4]byte{'U', 'R', 'D', 'L'}
	current_dir := 0

	walls := map[Pos]struct{}{}

	start := Pos{}
	pos := Pos{}

	for _, move := range moves {
		if move.dir == 'L' {
			if current_dir == 0 {
				current_dir = 3
			} else {
				current_dir--
			}
		} else if move.dir == 'R' {
			if current_dir == 3 {
				current_dir = 0
			} else {
				current_dir++
			}
		}

		for i := 0; i < move.amount; i++ {
			switch dirs[current_dir] {
			case 'U':
				pos.y--
			case 'R':
				pos.x++
			case 'D':
				pos.y++
			case 'L':
				pos.x--
			}
			walls[pos] = struct{}{}
		}
	}

	end := pos
	delete(walls, end)

	queue := []Entry{Entry{pos: start}}
	visited := map[Pos]struct{}{}
	visited[start] = struct{}{}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		pos := entry.pos

		if pos.x == end.x && pos.y == end.y {
			return entry.distance
		}

		for _, dir := range cardinal_dirs {
			new_pos := Pos{x: pos.x + dir[0], y: pos.y + dir[1]}
			if _, ok := walls[new_pos]; ok {
				continue
			}
			if _, ok := visited[new_pos]; ok {
				continue
			}
			queue = append(queue, Entry{pos: new_pos, distance: entry.distance + 1})
			visited[new_pos] = struct{}{}
		}
	}

	return -1
}

func crosses_wall(pos Pos, wall Wall) bool {
	if wall.vertical && wall.start.x == pos.x && pos.y >= wall.start.y && pos.y <= wall.end.y {
		return true
	}
	if !wall.vertical && wall.start.y == pos.y && pos.x >= wall.start.x && pos.x <= wall.end.x {
		return true
	}
	return false
}

func crosses_walls(pos Pos, walls []Wall) bool {
	for _, wall := range walls {
		if crosses_wall(pos, wall) {
			return true
		}
	}
	return false
}

func find_exit_large(moves []Move) int {
	dirs := [4]byte{'U', 'R', 'D', 'L'}
	current_dir := 0

	walls := []Wall{}

	x_turns := map[int]struct{}{}
	y_turns := map[int]struct{}{}

	x_turns[0] = struct{}{}
	y_turns[0] = struct{}{}

	start := Pos{}
	pos := Pos{}

	for _, move := range moves {
		if move.dir == 'L' {
			if current_dir == 0 {
				current_dir = 3
			} else {
				current_dir--
			}
		} else if move.dir == 'R' {
			if current_dir == 3 {
				current_dir = 0
			} else {
				current_dir++
			}
		}

		last_pos := pos
		switch dirs[current_dir] {
		case 'U':
			pos.y -= move.amount
		case 'R':
			pos.x += move.amount
		case 'D':
			pos.y += move.amount
		case 'L':
			pos.x -= move.amount
		}

		wall := Wall{start: last_pos, end: pos}
		if dirs[current_dir] == 'U' || dirs[current_dir] == 'D' {
			wall.vertical = true
		}
		if wall.start.x > wall.end.x || wall.start.y > wall.end.y {
			wall.start, wall.end = wall.end, wall.start
		}
		walls = append(walls, wall)

		x_turns[pos.x-1] = struct{}{}
		x_turns[pos.x] = struct{}{}
		x_turns[pos.x+1] = struct{}{}
		y_turns[pos.y-1] = struct{}{}
		y_turns[pos.y] = struct{}{}
		y_turns[pos.y+1] = struct{}{}
	}

	end := pos

	visited := map[Pos]struct{}{}

	queue := make(PriorityQueue, 1)
	queue[0] = &Entry{pos: start}
	heap.Init(&queue)

	for queue.Len() > 0 {
		entry := heap.Pop(&queue).(*Entry)

		pos := entry.pos

		if pos.x == end.x && pos.y == end.y {
			return entry.distance
		}

		for _, dir := range cardinal_dirs {
			dx, dy := dir[0], dir[1]

			new_pos := pos
			dist := 0

			for {
				new_pos.x += dx
				new_pos.y += dy
				dist++

				if new_pos.x == end.x && new_pos.y == end.y {
					return entry.distance + dist
				}

				if dy == 0 {
					if _, ok := x_turns[new_pos.x]; ok {
						if crosses_walls(new_pos, walls) {
							new_pos.x -= dx
							new_pos.y -= dy
							dist--
						}
						break
					}
				}
				if dx == 0 {
					if _, ok := y_turns[new_pos.y]; ok {
						if crosses_walls(new_pos, walls) {
							new_pos.x -= dx
							new_pos.y -= dy
							dist--
						}
						break
					}
				}
			}

			if _, ok := visited[new_pos]; !ok {
				new_dist := entry.distance + dist
				heap.Push(&queue, &Entry{pos: new_pos, distance: new_dist})
				visited[new_pos] = struct{}{}
			}
		}

	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 15, 1

	data := loader.GetOneLine()
	moves := parse_data(data)
	part1 := find_exit(moves)

	loader.Part = 2
	data = loader.GetOneLine()
	moves = parse_data(data)
	part2 := find_exit(moves)

	loader.Part = 3
	data = loader.GetOneLine()
	moves = parse_data(data)
	part3 := find_exit_large(moves)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
