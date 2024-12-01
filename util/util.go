package util

import (
	"bufio"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

func ReadInput() ([]string, error) {
	testmode := pflag.BoolP("test", "t", false, "use test.txt instead of input.txt")
	if *testmode {
		return ReadFromFile("test.txt")
	}

	return ReadFromFile("input.txt")
}

func ReadFromFile(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func ShouldRunPart1() bool {
	args := pflag.Args()
	if len(args) > 1 {
		panic("too many arguments")
	}

	if len(args) == 0 {
		return true
	}

	if args[0] == "part1" {
		return true
	}

	if args[0] == "part2" {
		return false
	}

	panic("invalid argument " + args[0])
}

func ShouldRunPart2() bool {
	return !ShouldRunPart1()
}
