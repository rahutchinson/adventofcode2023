package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Map struct {
	source      int64
	destination int64
	length      int64
}

func (m *Map) GetEndRangeSource() int64 {
	return m.source + m.length
}

func (m *Map) GetEndRangeDestination() int64 {
	return m.destination + m.length
}

type SeedMap struct {
	seed  int64
	soil  int64
	fert  int64
	water int64
	light int64
	temp  int64
	humid int64
	loc   int64
}

func trimAtoi(toTrim string) int64 {
	num, _ := strconv.Atoi(strings.TrimSpace(toTrim))
	return int64(num)

}

// inRange checks if a value is within a range of integers.
func inRange(value, start, end int64) bool {
	return value >= start && value <= end
}

func AbsInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func mapToDestination(m []Map, sourceVal int64) int64 {
	for _, ma := range m {
		if inRange(sourceVal, ma.source, ma.GetEndRangeSource()) {
			var diff int64
			var dest int64
			diff = ma.GetEndRangeSource() - sourceVal
			dest = ma.GetEndRangeDestination() - diff
			return dest

		}
	}
	return sourceVal
}

func main() {
	dat, err := os.ReadFile("./inputs/day5/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	var seeds []int64
	mapOfLists := make(map[string][]Map)
	currentMap := ""
	for _, line := range temp {
		split := strings.Split(line, ":")
		if split[0] == "seeds" {
			nums := strings.Split(strings.TrimSpace(split[1]), " ")
			for _, num := range nums {
				conv, _ := strconv.Atoi(num)
				seeds = append(seeds, int64(conv))
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
		}
	}

	minLocation := SeedMap{}
	minLocation.loc = 203685477580700
	for _, seed := range seeds {
		ss := SeedMap{}
		ss.seed = seed
		ss.soil = mapToDestination(mapOfLists["seed-to-soil map"], seed)
		ss.fert = mapToDestination(mapOfLists["soil-to-fertilizer map"], ss.soil)
		ss.water = mapToDestination(mapOfLists["fertilizer-to-water map"], ss.fert)
		ss.light = mapToDestination(mapOfLists["water-to-light map"], ss.water)
		ss.temp = mapToDestination(mapOfLists["light-to-temperature map"], ss.light)
		ss.humid = mapToDestination(mapOfLists["temperature-to-humidity map"], ss.temp)
		ss.loc = mapToDestination(mapOfLists["humidity-to-location map"], ss.humid)
		fmt.Println(ss)
		if ss.loc < minLocation.loc {
			minLocation = ss
		}
	}
	fmt.Println(minLocation.loc)

}
