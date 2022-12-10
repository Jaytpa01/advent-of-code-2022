package ten

import (
	"fmt"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func B() string {
	// input := utils.NewFileScanner("./10/sample1.txt")
	// input := utils.NewFileScanner("./10/sample2.txt")
	input := utils.NewFileScanner("./10/input.txt")
	defer input.Close()

	cycle := 1

	// map of cycle: addx value
	toAdd := map[int]int{}

	for input.Scan() {
		instruction := input.Text()

		if instruction == "noop" {
			cycle++
			continue
		}

		var addx int
		fmt.Sscanf(instruction, "addx %d", &addx)

		// after two cycles, the addx op will be completed
		toAdd[cycle+2] = addx
		cycle += 2
	}

	x := 1
	spriteRange := make([]int, 3)

	crt := ""
	for cycle := 1; cycle <= 240; cycle++ {
		v, ok := toAdd[cycle]
		if ok {
			x += v
		}

		spriteRange[0], spriteRange[1], spriteRange[2] = x-1, x, x+1

		if intIn((cycle%40)-1, spriteRange) {
			crt += "#"
		} else {
			crt += " "
		}

	}

	output := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", crt[:40], crt[40:80], crt[80:120], crt[120:160], crt[160:200], crt[200:240])

	return output
}

func intIn(i int, arr []int) bool {
	for _, t := range arr {
		if i == t {
			return true
		}
	}

	return false
}
