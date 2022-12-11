package eleven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func B() int {
	// input := utils.OpenInputFile("./11/sample.txt")
	input := utils.OpenInputFile("./11/input.txt")
	monkeyInputs := strings.Split(input, "\n\n")

	monkeys := make([]*monkey, len(monkeyInputs))
	for i := range monkeys {
		monkeys[i] = &monkey{}
	}

	mod := 1

	for i, monkeyInput := range monkeyInputs {
		lines := strings.Split(monkeyInput, "\n")

		// strip whitespace from all lines
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

		var operation string
		var val string
		fmt.Sscanf(lines[2], "Operation: new = old %s %s", &operation, &val)
		monkeys[i].calculation = calculation{operation, val}

		var test int
		fmt.Sscanf(lines[3], "Test: divisible by %d", &test)
		monkeys[i].divisibleBy = test
		mod *= test

		ifTrue, ifFalse := 0, 0
		fmt.Sscanf(lines[4], "If true: throw to monkey %d", &ifTrue)
		fmt.Sscanf(lines[5], "If false: throw to monkey %d", &ifFalse)
		monkeys[i].ifTrue = ifTrue
		monkeys[i].ifFalse = ifFalse
	}

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {

			numOfItems := len(monkey.items)
			for i := 0; i < numOfItems; i++ {
				monkey.inspections++

				item := monkey.items.dequeue()

				item = monkey.performOperation(item)

				item %= mod

				if item%monkey.divisibleBy == 0 {
					monkeys[monkey.ifTrue].items.enqueue(item)
				} else {
					monkeys[monkey.ifFalse].items.enqueue(item)
				}

			}

		}

	}

	inspections := make([]int, len(monkeys))
	for i := range inspections {
		inspections[i] = monkeys[i].inspections
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Println(inspections)

	return inspections[0] * inspections[1]
}
