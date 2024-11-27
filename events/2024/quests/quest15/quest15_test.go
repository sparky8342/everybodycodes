package quest15

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"#####.#####",
		"#.........#",
		"#.######.##",
		"#.........#",
		"###.#.#####",
		"#H.......H#",
		"###########",
	}
	got := simple_solve(parse_data(data))
	var want uint16 = 26

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"##########.##########",
		"#...................#",
		"#.###.##.###.##.#.#.#",
		"#..A#.#..~~~....#A#.#",
		"#.#...#.~~~~~...#.#.#",
		"#.#.#.#.~~~~~.#.#.#.#",
		"#...#.#.B~~~B.#.#...#",
		"#...#....BBB..#....##",
		"#C............#....C#",
		"#####################",
	}

	got := simple_solve(parse_data(data))
	var want uint16 = 38

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
