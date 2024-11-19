package loader

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Event string
var Quest int
var Part int

func get_filename() string {
	return fmt.Sprintf("inputs/%s/everybody_codes_e%s_q%02d_p%d.txt", Event, Event, Quest, Part)
}

func GetOneLine() []byte {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return data
}

func GetStrings() []string {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}
