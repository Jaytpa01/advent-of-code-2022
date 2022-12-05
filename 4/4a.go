package four

import (
	"strconv"
	"strings"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func A() int {
	// scanner := utils.NewFileScanner("./4/sample.txt")
	scanner := utils.NewFileScanner("./4/input.txt")
	defer scanner.Close()

	total := 0

	for scanner.Scan() {
		input := scanner.Text()

		splitInput := strings.Split(input, ",")

		a, b := strings.Split(splitInput[0], "-"), strings.Split(splitInput[1], "-")

		if RangeIsWithin(a, b) || RangeIsWithin(b, a) {
			total++
		}

	}
	return total
}

func RangeIsWithin(a, b []string) bool {
	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])
	return a0 <= b0 && a1 >= b1
}
