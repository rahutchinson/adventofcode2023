package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// up (-1,0), etc
type Direction struct {
	v int
	h int
}

func direction(pipe rune) Direction {
	switch pipe {
	case 'F':

	}
	return Direction{0, 1}
}

type Pair struct {
	X int
	Y int
}

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/
func traversePipe(loc Pair, ar []string, maps map[Pair]string) int {
	if ar[loc.Y][loc.X] == 'S' {
		return 1
	}
	if loc.Y != len(ar)-1 {
		switch {
		//going down
		case loc.Y != len(ar[loc.Y])-2 && ar[loc.Y-1] == "|":
			maps[Pair{loc.Y - 1, loc.X}] = "|"
			return traversePipe(Pair{loc.Y - 1, loc.X}, ar, maps) + 1
		case loc.Y != len(ar[loc.Y])-1 && ar[loc.Y-1] == "L":
			maps[Pair{loc.Y - 1, loc.X}] = "L"
			return traversePipe(Pair{loc.Y - 1, loc.X}, ar, maps) + 1
		case loc.Y != len(ar[loc.Y])-1 && ar[loc.Y-1] == "J":
			maps[Pair{loc.Y - 1, loc.X}] = "J"
			return traversePipe(Pair{loc.Y - 1, loc.X}, ar, maps) + 1
			//going up
		case loc.Y != len(ar[loc.Y])+2 && ar[loc.Y+1] == "|":
			maps[Pair{loc.Y - 1, loc.X}] = "J"
			return traversePipe(Pair{loc.Y - 1, loc.X}, ar, maps) + 1
		case loc.Y != len(ar[loc.Y])+1 && ar[loc.Y+1] == "|":
			maps[Pair{loc.Y - 1, loc.X}] = "J"
			return traversePipe(Pair{loc.Y - 1, loc.X}, ar, maps) + 1

		}

	}

}

func main() {
	dat, err := os.ReadFile("./inputs/day10/example.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	pipes := make([]string, 0, len(temp))
	sLoc := Pair{0, 0}
	for y, line := range temp {
		pipes = append(pipes, line)
		s := strings.Index(line, "S")
		if s != -1 {
			sLoc = Pair{s, y}
		}
	}
	fmt.Println(sLoc)
	fmt.Println(pipes)
}
