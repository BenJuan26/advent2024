package main

var part1Contents = []byte(`package main

import (
	"fmt"

	"github.com/BenJuan26/advent2024/util"
)

func Part1() {
	lines, err := util.ReadInput()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		line += " " // do something
	}

	fmt.Println("answer")
}

func main() {
	if util.ShouldRunPart1() {
		Part1()
	}

	if util.ShouldRunPart2() {
		Part2()
	}
}
`)

var part2Contents = []byte(`package main

import (
	"fmt"

	"github.com/BenJuan26/advent2024/util"
)

func Part2() {
	lines, err := util.ReadInput()
	if err != nil {
		panic(err)
	}

	fmt.Println(lines[0])
}
`)
