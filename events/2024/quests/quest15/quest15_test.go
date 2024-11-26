package quest15

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"#####.#####",
		"#.........#",
		"#.######.##",
		"#.........#",
		"###.#.#####",
		"#H.......H#",
		"###########",
	}
	got := find_path(grid)
	want := 26

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
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

	got := find_path(grid)
	want := 38

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
