package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Map struct {
	source      int
	destination int
	length      int
}

func (m *Map) GetEndRangeSource() int {
	return m.source + m.length
}

func (m *Map) GetEndRangeDestination() int {
	return m.destination + m.length
}

type SeedMap struct {
	seed  int
	soil  int
	fert  int
	water int
	light int
	temp  int
	humid int
	loc   int
}

func trimAtoi(toTrim string) int {
	num, _ := strconv.Atoi(strings.TrimSpace(toTrim))
	return int(num)

}

// inRange checks if a value is within a range of integers.
func inRange(value, start, end int) bool {
	return value >= start && value <= end
}

func Absint(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func mapToDestination(m []Map, sourceVal int) int {
	for _, ma := range m {
		if inRange(sourceVal, ma.source, ma.GetEndRangeSource()) {
			var diff int
			var dest int
			diff = ma.GetEndRangeSource() - sourceVal
			dest = ma.GetEndRangeDestination() - diff
			return dest
		}
	}
	return sourceVal
}

func main() {
	startTime := time.Now()
	dat, err := os.ReadFile("./inputs/day5/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	var seeds []int
	mapOfLists := make(map[string][]Map)
	currentMap := ""
	for _, line := range temp {
		split := strings.Split(line, ":")
		if split[0] == "seeds" {
			nums := strings.Split(strings.TrimSpace(split[1]), " ")
			for _, num := range nums {
				conv, _ := strconv.Atoi(num)
				seeds = append(seeds, int(conv))
			}
			continue
		}
		if line == "" {
			continue
		}
		if unicode.IsLetter(rune(line[0])) {
			currentMap = split[0]
			continue
		}
		if unicode.IsNumber(rune(line[0])) {
			nums := strings.Split(strings.TrimSpace(line), " ")
			mapOfLists[currentMap] = append(mapOfLists[currentMap], Map{trimAtoi(nums[1]), trimAtoi(nums[0]), trimAtoi(nums[2])})
			continue
		}
	}

	minLocation := SeedMap{}
	minLocation.loc = 203685477580700
	for i := 0; i < len(seeds)-1; i += 2 {
		for j := 0; j < int(seeds[i+1]); j++ {
			ss := SeedMap{}
			ss.seed = seeds[i] + int(j)
			ss.soil = mapToDestination(mapOfLists["seed-to-soil map"], ss.seed)
			ss.fert = mapToDestination(mapOfLists["soil-to-fertilizer map"], ss.soil)
			ss.water = mapToDestination(mapOfLists["fertilizer-to-water map"], ss.fert)
			ss.light = mapToDestination(mapOfLists["water-to-light map"], ss.water)
			ss.temp = mapToDestination(mapOfLists["light-to-temperature map"], ss.light)
			ss.humid = mapToDestination(mapOfLists["temperature-to-humidity map"], ss.temp)
			ss.loc = mapToDestination(mapOfLists["humidity-to-location map"], ss.humid)
			if ss.loc < minLocation.loc {
				minLocation = ss
			}
		}
		fmt.Println(minLocation)
	}
	duration := time.Since(startTime)
	fmt.Println("Execution time:", duration)
}
