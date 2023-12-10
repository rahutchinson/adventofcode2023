package main

import (
	"fmt"
	"math"
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

func traversePipes(grid [][]rune, startX, startY int, path *[][]int) {
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

		// Append current position to path
		*path = append(*path, []int{x, y})

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

func stringsToRunes(strings []string) [][]rune {
	var runes [][]rune
	for _, str := range strings {
		runes = append(runes, []rune(str))
	}
	return runes
}

func shoelace(vertices [][]int) float64 {
	area := 0.0
	n := len(vertices)

	for i := 0; i < n-1; i++ {
		area += float64(vertices[i][0]*vertices[i+1][1] - vertices[i+1][0]*vertices[i][1])
	}

	// Closing the polygon
	area += float64(vertices[n-1][0]*vertices[0][1] - vertices[0][0]*vertices[n-1][1])

	return math.Abs(area) / 2.0
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
	path := make([][]int, 0)
	pipes_r := stringsToRunes(pipes)
	traversePipes(pipes_r, sLoc.X, sLoc.Y, &path)
	// Prepare a visited grid
	visited := make([][]bool, len(pipes_r))
	for i := range visited {
		visited[i] = make([]bool, len(pipes_r[i]))
	}
	// Find a reliable starting point for flood fill
	// This is a placeholder; you'll need a robust method based on your grid and path
	area := shoelace(path)
	// area = numIneed + len(path)/2 -1
	// area + 1 - len(path)/2 = numIneed 
	paths := len(path)/2
	fmt.Println(int(area) + 1 - int(paths))

}

// 19600 too high
// 985 too high
// 513 too high
