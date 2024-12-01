package main

import "fmt"

func Part2() {
	list1, list2 := SortedLists()

	score := 0
	j := 0
	memo := map[int]int{}
	for _, num1 := range list1 {
		if increase, found := memo[num1]; found {
			score += increase
			continue
		}

		for j < len(list2) && list2[j] < num1 {
			j += 1
		}

		freq := 0
		for j < len(list2) && list2[j] == num1 {
			freq += 1
			j += 1
		}

		scoreIncrease := num1 * freq
		memo[num1] = scoreIncrease
		score += scoreIncrease
	}

	fmt.Println(score)
}
