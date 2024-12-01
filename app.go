package main

import (
	"fmt"
	"os"
	"path"

	"github.com/BenJuan26/advent2024/util"
	"github.com/spf13/pflag"
)

func LatestDay() int {
	max := 0

	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if entry.Name()[:3] != "day" {
			continue
		}

		dayNum := util.MustAtoi(entry.Name()[3:])
		if dayNum > max {
			max = dayNum
		}
	}

	return max
}

func InitFile(filename string, contents []byte) {
	// file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	panic(err)
	// }

	// if contents != "" {

	// }

	// err = file.Close()
	// if err != nil {
	// 	panic(err)
	// }

	err := os.WriteFile(filename, contents, 0666)
	if err != nil {
		panic(err)
	}
}

func AddDay() {
	newNum := LatestDay() + 1
	dirName := fmt.Sprintf("day%02d", newNum)

	err := os.Mkdir(dirName, 0777)
	if err != nil {
		panic(err)
	}

	InitFile(path.Join(dirName, "input.txt"), []byte{})
	InitFile(path.Join(dirName, "test.txt"), []byte{})
	InitFile(path.Join(dirName, "part1.go"), part1Contents)
	InitFile(path.Join(dirName, "part2.go"), part2Contents)
}

func main() {
	pflag.Parse()
	args := pflag.Args()
	if len(args) < 1 {
		panic("no command provided")
	}

	if len(args) > 1 {
		panic(fmt.Sprintf("expected 1 command, got %d", len(args)))
	}

	switch args[0] {
	case "addday":
		AddDay()
	default:
		panic(fmt.Sprintf("invalid command: %s", args[0]))
	}
}
