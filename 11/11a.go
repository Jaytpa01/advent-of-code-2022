package eleven

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func A() int {
	input := utils.OpenInputFile("./11/sample.txt")
	// input := utils.OpenInputFile("./11/sample.txt")
	monkeyInputs := strings.Split(input, "\n\n")

	monkeys := make([]monkey, len(monkeyInputs))

	for i, monkeyInput := range monkeyInputs {
		lines := strings.Split(monkeyInput, "\n")

		for i, line := range lines {
			line = strings.TrimSpace(line)
			lines[i] = line
		}

		// convert starting items input to array of ints
		startingItems := lines[1][len("Starting items: "):]
		itemStrArr := strings.Split(startingItems, ", ")
		for _, str := range itemStrArr {
			item, _ := strconv.Atoi(str)
			monkeys[i].items = append(monkeys[i].items, item)
		}

	}

	fmt.Println(monkeys)
	return 0
}
