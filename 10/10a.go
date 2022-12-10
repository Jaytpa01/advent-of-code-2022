package ten

import (
	"fmt"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

var cyclesToCheck map[int]bool = map[int]bool{
	20:  true,
	60:  true,
	100: true,
	140: true,
	180: true,
	220: true,
}

func A() int {
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
	signalStrength := 0
	for i := 1; i < 222; i++ {
		v, ok := toAdd[i]
		if ok {
			x += v
		}

		if cyclesToCheck[i] {
			signalStrength += i * x
		}
	}

	fmt.Println(x)

	return signalStrength
}
