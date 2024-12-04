package quest20

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"#....S....#",
		"#.........#",
		"#---------#",
		"#.........#",
		"#..+.+.+..#",
		"#.+-.+.++.#",
		"#.........#",
	}

	got := highest_altitude(grid)
	want := 1045

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"####S####",
		"#-.+++.-#",
		"#.+.+.+.#",
		"#-.+.+.-#",
		"#A+.-.+C#",
		"#.+-.-+.#",
		"#.+.B.+.#",
		"#########",
	}

	got := shortest_path(grid)
	want := 24

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	grid := []string{
		"###############S###############",
		"#+#..-.+.-++.-.+.--+.#+.#++..+#",
		"#-+-.+-..--..-+++.+-+.#+.-+.+.#",
		"#---.--+.--..++++++..+.-.#.-..#",
		"#+-+.#+-.#-..+#.--.--.....-..##",
		"#..+..-+-.-+.++..-+..+#-.--..-#",
		"#.--.A.-#-+-.-++++....+..C-...#",
		"#++...-..+-.+-..+#--..-.-+..-.#",
		"#..-#-#---..+....#+#-.-.-.-+.-#",
		"#.-+.#+++.-...+.+-.-..+-++..-.#",
		"##-+.+--.#.++--...-+.+-#-+---.#",
		"#.-.#+...#----...+-.++-+-.+#..#",
		"#.---#--++#.++.+-+.#.--..-.+#+#",
		"#+.+.+.+.#.---#+..+-..#-...---#",
		"#-#.-+##+-#.--#-.-......-#..-##",
		"#...+.-+..##+..+B.+.#-+-++..--#",
		"###############################",
	}

	got := shortest_path(grid)
	want := 78

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	grid := []string{
		"###############S###############",
		"#-----------------------------#",
		"#-------------+++-------------#",
		"#-------------+++-------------#",
		"#-------------+++-------------#",
		"#-----------------------------#",
		"#-----------------------------#",
		"#-----------------------------#",
		"#--A-----------------------C--#",
		"#-----------------------------#",
		"#-----------------------------#",
		"#-----------------------------#",
		"#-----------------------------#",
		"#-----------------------------#",
		"#-----------------------------#",
		"#--------------B--------------#",
		"#-----------------------------#",
		"#-----------------------------#",
		"###############################",
	}

	got := shortest_path(grid)
	want := 206

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
