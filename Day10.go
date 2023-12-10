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
	Y int
	X int
}

func traversePipes(grid [][]rune, startX, startY int) {
	x, y := startX, startY
	direction := "east" // assuming the animal starts moving east
	started := false    // flag to indicate if the traversal has started
	running_count := 0
	for {
		running_count++
		if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
			fmt.Println("Exited the grid.")
			break
		}

		switch grid[y][x] {
		case '|':
			if direction != "north" && direction != "south" {
				fmt.Println("Reached an invalid pipe for current direction.")
				return
			}
		case '-':
			if direction != "east" && direction != "west" {
				fmt.Println("Reached an invalid pipe for current direction.")
				return
			}
		case 'L':
			if direction == "south" {
				direction = "east"
			} else if direction == "west" {
				direction = "north"
			} else {
				fmt.Println("Reached an invalid pipe for current direction.")
				return
			}
		case 'J':
			if direction == "south" {
				direction = "west"
			} else if direction == "east" {
				direction = "north"
			} else {
				fmt.Println("Reached an invalid pipe for current direction.")
				return
			}
		case '7':
			if direction == "north" {
				direction = "west"
			} else if direction == "east" {
				direction = "south"
			} else {
				fmt.Println("Reached an invalid pipe for current direction.")
				return
			}
		case 'F':
			if direction == "north" {
				direction = "east"
			} else if direction == "west" {
				direction = "south"
			} else {
				fmt.Println("Reached an invalid pipe for current direction.")
				return
			}
		case '.':
			fmt.Println("Reached the end of the pipe.")
			return
		case 'S':
			if started {
				fmt.Println((running_count - 1) / 2)
				fmt.Println("Returned to the start position.")
				return
			}
			started = true
		default:
			fmt.Println("Encountered an unknown character.")
			return
		}

		// Move to the next tile based on the current direction
		switch direction {
		case "north":
			y--
		case "south":
			y++
		case "east":
			x++
		case "west":
			x--
		}

		fmt.Printf("Current position: (%d, %d), moving %s\n", x, y, direction)
	}
}
func isInside(grid [][]rune, x, y int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y])
}

func floodFill(grid [][]rune, x, y int) int {
	if !isInside(grid, x, y) || grid[y][x] != '.' {
		return 0
	}

	// Mark the current cell as visited
	grid[y][x] = 'x'

	// Recursively call for all directions
	count := 1
	count += floodFill(grid, x+1, y)
	count += floodFill(grid, x-1, y)
	count += floodFill(grid, x, y+1)
	count += floodFill(grid, x, y-1)

	return count
}

func findStartInsideLoop(grid [][]rune, path [][]int) (int, int) {
	// Simplified: Find a point next to the first point in the path.
	// In a more complex loop, you would need to find a point that is guaranteed to be inside the loop.
	for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		x := path[0][0] + dir[0]
		y := path[0][1] + dir[1]
		if isInside(grid, x, y) && grid[y][x] == '.' {
			return x, y
		}
	}
	return -1, -1
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
func traversePipe(loc, last Pair, ar []string, maps map[Pair]string) int {
	if ar[loc.Y][loc.X] == 'S' {
		fmt.Println(maps)
		return 1
	}
	fmt.Println(maps)
	fmt.Println(string(ar[loc.Y][loc.X]))
	val := string(ar[loc.Y][loc.X])
	switch {
	//going up
	case loc.Y < len(ar[loc.Y])-1 && ar[loc.Y-1][loc.X] == '|' && last.Y != loc.Y-1:
		maps[Pair{loc.Y - 1, loc.X}] = "|"
		return traversePipe(Pair{loc.Y - 1, loc.X}, loc, ar, maps) + 1
	case loc.Y != len(ar[loc.Y])-1 && ar[loc.Y-1][loc.X] == 'L' && last.Y != loc.Y-1:
		maps[Pair{loc.Y - 1, loc.X}] = "L"
		return traversePipe(Pair{loc.Y - 1, loc.X}, loc, ar, maps) + 1
	case loc.Y != len(ar[loc.Y])-1 && ar[loc.Y-1][loc.X] == 'J' && last.Y != loc.Y-1:
		maps[Pair{loc.Y - 1, loc.X}] = "J"
		return traversePipe(Pair{loc.Y - 1, loc.X}, loc, ar, maps) + 1
	//going down
	case val != "J" && loc.Y != len(ar[loc.Y])+2 && ar[loc.Y+1][loc.X] == '|' && last.Y != loc.Y+1:
		maps[Pair{loc.Y + 1, loc.X}] = "|"
		return traversePipe(Pair{loc.Y + 1, loc.X}, loc, ar, maps) + 1
	case val != "J" && loc.Y != len(ar[loc.Y])+1 && ar[loc.Y+1][loc.X] == '7' && last.Y != loc.Y+1:
		maps[Pair{loc.Y + 1, loc.X}] = "7"
		return traversePipe(Pair{loc.Y + 1, loc.X}, loc, ar, maps) + 1
	case val != "J" && loc.Y != len(ar[loc.Y])+1 && ar[loc.Y+1][loc.X] == 'F' && last.Y != loc.Y+1:
		maps[Pair{loc.Y + 1, loc.X}] = "F"
		return traversePipe(Pair{loc.Y + 1, loc.X}, loc, ar, maps) + 1
		// going right
	case loc.X < len(ar[loc.Y])-1 && ar[loc.Y][loc.X+1] == '-' && last.X != loc.X+1:
		maps[Pair{loc.Y, loc.X + 1}] = "-"
		return traversePipe(Pair{loc.Y, loc.X + 1}, loc, ar, maps) + 1
		//going left
	case loc.X > 0 && ar[loc.Y][loc.X-1] == '-' && last.X != loc.X-1:
		maps[Pair{loc.Y, loc.X - 1}] = "-"
		return traversePipe(Pair{loc.Y, loc.X - 1}, loc, ar, maps) + 1
	}
	return 0
}

func stringsToRunes(strings []string) [][]rune {
	var runes [][]rune
	for _, str := range strings {
		runes = append(runes, []rune(str))
	}
	return runes
}

func main() {
	dat, err := os.ReadFile("./inputs/day10/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	pipes := make([]string, 0, len(temp))
	sLoc := Pair{0, 0}
	for y, line := range temp {
		pipes = append(pipes, line)
		s := strings.Index(line, "S")
		if s != -1 {
			sLoc = Pair{y, s}
		}
	}

	// length := traversePipe(Pair{sLoc.Y, sLoc.X + 1}, sLoc, pipes, make(map[Pair]string))
	traversePipes(stringsToRunes(pipes), sLoc.X, sLoc.Y)
	fmt.Println(sLoc)
	// fmt.Println(length)
}
