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

type Pair struct {
	X int
	Y int
}

func createCol(uni []string, col int) string {
	cols := ""
	for _, val := range uni {
		cols = cols + string(val[col])
	}
	return cols
}

func expand(uni []string) []string {
	empty_cols := make([]int, 0, len(uni[0]))
	empty_rows := make([]int, 0, len(uni))

	for row, val := range uni {
		if val == strings.Repeat(".", len(val)) {
			empty_rows = append(empty_rows, row)
		}
	}
	for i := 0; i < len(uni[0]); i++ {
		if createCol(uni, i) == strings.Repeat(".", len(uni)) {
			empty_cols = append(empty_cols, i)
		}
	}
	newTable := uni
	val_to_expand := 10
	for val, row := range empty_rows {
		// new row to add
		newRow := strings.Repeat(".", len(newTable[0]))
		newTableRow := make([]string, len(newTable)+val_to_expand)

		copy(newTableRow, newTable[:row+val])
		for i := 1; i < val_to_expand; i++ {
			newTableRow[row+val+i] = newRow
		}
		copy(newTableRow[row+val+val_to_expand:], newTable[row+val:])

		newTable = newTableRow
	}
	for val, row := range newTable {
		newRow := row
		for num, col := range empty_cols {
			newRow = newRow[:col+num] + strings.Repeat(".", val_to_expand) + newRow[col+num:]
		}
		newTable[val] = newRow
	}
	return newTable
}

func printUni(uni []string) {
	for i := 0; i < len(uni); i++ {
		fmt.Println(uni[i])
	}
}

func mapGalxies(uni []string) map[int]Pair {
	r := make(map[int]Pair)
	running_count := 1
	for y, line := range uni {
		for x, char := range line {
			if char == '#' {
				r[running_count] = Pair{x, y}
				running_count++
			}
		}
	}
	return r
}

func Absint(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distance2D(a, b Pair) int {
	return Absint(b.X-a.X) + Absint(b.Y-a.Y)
}

//first expand the galxy
// find all the galaxy locations
// find the distance between each galaxy

func main() {
	dat, err := os.ReadFile("./inputs/day11/example.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	expanded := expand(temp)
	mapUni := mapGalxies(expanded)
	running_sum := 0
	for g := 1; g < len(mapUni); g++ {
		min := 10000000
		for i := g; i < len(mapUni)+1; i++ {
			if g != i && distance2D(mapUni[g], mapUni[i]) < min {
				running_sum += distance2D(mapUni[g], mapUni[i])
			}
		}
		fmt.Printf("num: %d min: %d \n", g, running_sum)
	}
	fmt.Println(running_sum)
}
