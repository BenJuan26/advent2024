package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BenJuan26/advent2024/util"
)

func SortedLists() ([]int, []int) {
	lines, err := util.ReadInput()
	if err != nil {
		panic(err)
	}

	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)
	for _, line := range lines {
		fields := strings.Split(line, "   ")

		num1 := util.MustAtoi(fields[0])
		list1 = append(list1, num1)

		num2 := util.MustAtoi(fields[1])
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return list1, list2
}

func Part1() {
	total := 0
	list1, list2 := SortedLists()
	for i := range list1 {
		diff := util.Abs(list2[i] - list1[i])
		total += diff
	}

	fmt.Println(total)
}

func main() {
	if util.ShouldRunPart1() {
		Part1()
	}

	if util.ShouldRunPart2() {
		Part2()
	}
}
