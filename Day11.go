package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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

func createRows(numRows, dots int) []string {
	r := make([]string, numRows)
	dot_s := strings.Repeat(".", dots)
	for i := 0; i < numRows; i++ {
		r[i] = dot_s
	}
	return r
}

func expandRows(uni []string, eVal int) []string {
	var newUni []string
	for _, row := range uni {
		newUni = append(newUni, row)
		if row == strings.Repeat(".", len(row)) {
			for i := 0; i < eVal; i++ {
				newUni = append(newUni, row)
			}
		}
	}
	return newUni
}

// this works
func expand2(uni []string, eVal int) (int, []string) {
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

	oldTable := expandRows(uni, eVal)

	for val, row := range oldTable {
		newRow := row
		for num, col := range empty_cols {
			num = num * eVal
			newRow = newRow[:col+num] + strings.Repeat(".", eVal) + newRow[col+num:]
		}
		oldTable[val] = newRow
	}
	return len(empty_rows), oldTable
}

func printUni(uni []string) {
	for i := 0; i < len(uni); i++ {
		fmt.Println(uni[i])
	}
	fmt.Println()
}

// given a universe - find the galaxies and return the labled galaxy with the location
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

func distance2DXY(a, b Pair) Pair {
	return Pair{Absint(b.X - a.X), Absint(b.Y - a.Y)}
}

//first expand the galxy
// find all the galaxy locations
// find the distance between each galaxy

type Gals struct {
	gal1 int
	gal2 int
}

func main() {
	startTime := time.Now()
	dat, err := os.ReadFile("./inputs/day11/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	// expanded := expand2(temp, 1)
	mapUni := mapGalxies(temp)
	// galaxy pair with X and Y diffs
	diff_map := make(map[Gals]Pair)
	for g := 1; g < len(mapUni)+1; g++ {
		for i := g; i < len(mapUni)+1; i++ {
			if g != i {
				diff_map[Gals{i, g}] = distance2DXY(mapUni[g], mapUni[i])
			}
		}
	}

	_, expanded10 := expand2(temp, 1)
	mapUni10 := mapGalxies(expanded10)
	diff_map10 := make(map[Gals]Pair)
	for g := 1; g < len(mapUni)+1; g++ {
		for i := g; i < len(mapUni)+1; i++ {
			if g != i {
				diff_map10[Gals{i, g}] = distance2DXY(mapUni10[g], mapUni10[i])
			}
		}
	}
	type isDiff struct {
		X int
		Y int
	}
	diff_impact := make(map[Gals]isDiff)
	for key, _ := range diff_map {
		diff_impact[key] = isDiff{diff_map10[key].X - diff_map[key].X, diff_map10[key].Y - diff_map[key].Y}
	}
	sum := 0
	for key, val := range diff_impact {
		sum += diff_map[key].X + val.X*999999 + diff_map[key].Y + val.Y*999999
	}
	fmt.Println(sum)
	duration := time.Since(startTime)
	fmt.Println("Execution time:", duration)
}

// 6108772 too low
// 678719321572 too low
// 678728799136 too low
// 678728808158
// 678728799116
