# advent2024

Advent of Code 2024

## Workflow

There are some helpful functions in the `util` package. To add a new day, just run from the top level:

```
go run . add
```

In the day's directory put the input from AoC into `input.txt`. To run part 1, run

```
go run .
```

or

```
go run . part1
```

To run part 2, run

```
go run . part2
```

For debugging, put whatever you want in `test.txt` and add the `-t` flag. The test input will be used instead of the main input.
