package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeIntList(list []string) []int {
	var int_list []int
	for _, str := range list {
		val, err := strconv.Atoi(string(str))
		if err == nil {
			int_list = append(int_list, val)
		}
	}
	return int_list
}

func calculateMinToBeat(time, distance int) int {
	for i := 1; i < time; i++ {
		time_left := time - i
		if i*time_left > distance {
			return i
		}
	}
	fmt.Println("can't beat record? min")
	return time
}

func calculateMaxToBeat(time, distance int) int {
	for i := time - 1; i > 0; i-- {
		time_left := time - i
		if i*time_left > distance {
			return i + 1
		}
	}
	fmt.Println("can't beat record? max")
	return time
}

func main() {
	dat, err := os.ReadFile("./inputs/day6/input.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")

	running_total := 1
	time_int := makeIntList(strings.Split(strings.Split(temp[0], ":")[1], " "))
	distance_int := makeIntList(strings.Split(strings.Split(temp[1], ":")[1], " "))
	for race, value := range time_int {
		total_to_beat := calculateMaxToBeat(value, distance_int[race]) - calculateMinToBeat(value, distance_int[race])
		fmt.Println(total_to_beat)
		running_total *= total_to_beat

	}
	fmt.Println(running_total)
}
