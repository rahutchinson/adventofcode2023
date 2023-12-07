package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Calculate base values for each hand type
const (
	FiveOfAKindBase  = 10000000
	FourOfAKindBase  = 1000000
	FullHouseBase    = 100000
	ThreeOfAKindBase = 10000
	TwoPairBase      = 1000
	OnePairBase      = 100
	NoPair           = 10
)

func sortString(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)

	// Sort the slice of runes
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Convert the slice of runes back to a string
	return string(runes)
}

// cardValue maps each card to its numerical value.
func cardValue(card rune) int {
	switch card {
	case '2', '3', '4', '5', '6', '7', '8', '9':
		return int(card - '0')
	case 'T':
		return 10
	case 'J':
		return 0
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return 0
	}
}

func removeJoker(slice []int) []int {
	for i, v := range slice {
		if v == 0 {
			// Remove the element at index i
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice // return the original slice if val is not found
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// evaluateHand returns a numerical value representing the strength of the hand.
func evaluateHand(hand string) int {
	frequencies := make(map[int]int)
	values := make([]int, 0, len(hand))

	// Count the frequency and store the values of each card label
	for _, card := range hand {
		val := cardValue(card)
		frequencies[val]++
	}

	for key, _ := range frequencies {
		values = append(values, key)
	}

	// Sort the values in descending order
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	// Evaluate the hand
	switch {
	case len(frequencies) == 1:
		return FiveOfAKindBase
	case len(frequencies) == 2:
		// we have a joker
		if contains(values, 0) {
			if frequencies[values[0]] == 1 && frequencies[values[1]] == 4 {
				return FiveOfAKindBase
			}
			if frequencies[values[1]] == 1 && frequencies[values[0]] == 4 {
				return FiveOfAKindBase
			}
			if frequencies[values[0]] == 2 && frequencies[values[1]] == 3 {
				return FiveOfAKindBase
			}
			if frequencies[values[1]] == 2 && frequencies[values[0]] == 3 {
				return FiveOfAKindBase
			}
		} else {
			if frequencies[values[0]] == 1 && frequencies[values[1]] == 4 {
				return FourOfAKindBase
			}
			if frequencies[values[1]] == 1 && frequencies[values[0]] == 4 {
				return FourOfAKindBase
			}
			if frequencies[values[0]] == 2 && frequencies[values[1]] == 3 {
				return FullHouseBase
			}
			if frequencies[values[1]] == 2 && frequencies[values[0]] == 3 {
				return FullHouseBase
			}
		}
		return 0
	case len(frequencies) == 3:
		// 0is1, 1is2, 2 is 2
		if contains(values, 0) {
			if frequencies[0] == 1 {
				values = removeJoker(values)
				if frequencies[values[0]] == 3 || frequencies[values[1]] == 3 {
					return FourOfAKindBase
				} else {
					return FullHouseBase
				}
			}
			if frequencies[0] == 3 {
				return FourOfAKindBase
			}
			if frequencies[0] == 2 {
				return FourOfAKindBase
			}
		} else {
			if frequencies[values[0]] == 1 && frequencies[values[1]] == 2 {
				if values[1] > values[2] {
					return TwoPairBase
				} else {
					return TwoPairBase
				}
			}
			// 0 is 2, 1is1, 2is2
			if frequencies[values[1]] == 1 && frequencies[values[2]] == 2 {
				if values[0] > values[2] {
					return TwoPairBase
				} else {
					return TwoPairBase
				}
			}
			// 0 is 2, 1is2, 2is1
			if frequencies[values[2]] == 1 && frequencies[values[0]] == 2 {
				if values[0] > values[1] {
					return TwoPairBase
				} else {
					return TwoPairBase
				}
			}

			//not a two pair
			return ThreeOfAKindBase
		}
		return 0
	case len(frequencies) == 4:
		if contains(values, 0) {
			return ThreeOfAKindBase
		}
		return OnePairBase
	default:
		if contains(values, 0) {
			return OnePairBase
		}
		// High card: use the highest value card
		return NoPair
	}
}

type keyValue struct {
	Strength int
	Bid      int
	Hand     string
}

func compareStrings(s1, s2 string, index int) string {

	if cardValue(rune(s1[index])) > cardValue(rune(s2[index])) {
		return s2
	} else if cardValue(rune(s1[index])) < cardValue(rune(s2[index])) {
		return s1
	}

	// If runes are equal, recurse with the next character
	return compareStrings(s1, s2, index+1)
}

func main() {
	dat, err := os.ReadFile("./inputs/day7/input.txt")
	// dat, err := os.ReadFile("./inputs/day7/example.txt")
	check(err)
	temp := strings.Split(string(dat), "\n")
	pairs := make([]keyValue, 0, 1000)
	for _, val := range temp {
		hand := strings.Split(strings.TrimSpace(val), " ")[0]
		bid, _ := strconv.Atoi(strings.Split(strings.TrimSpace(val), " ")[1])
		// fmt.Printf("hand %s score %d\n", hand, evaluateHand(hand))
		pairs = append(pairs, keyValue{Strength: evaluateHand(hand), Bid: bid, Hand: hand})
	}

	// Sort the slice by value in descending order
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Strength == pairs[j].Strength {
			return compareStrings(pairs[i].Hand, pairs[j].Hand, 0) == pairs[i].Hand
		}
		return pairs[i].Strength < pairs[j].Strength
	})

	total_sum := 0
	// Process the keys in sorted order
	for rank, pair := range pairs {
		fmt.Printf("hand %s rank %d strength %d bid %d\n", pair.Hand, rank, pair.Strength, pair.Bid)

		total_sum += (rank + 1) * pair.Bid
	}
	fmt.Println(total_sum)

	fmt.Println(cardValue('7'))
}

//.        253933213
//.        253640380
//.        253969937 -- not right
//		   253970704 -- not right
//.        254035226 --
//too high 254001575
