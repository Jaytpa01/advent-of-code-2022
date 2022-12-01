package one

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
--- Part Two ---
By the time you calculate the answer to the Elves' question, they've already realized that the Elf carrying the most Calories of food might eventually run out of snacks.

To avoid this unacceptable situation, the Elves would instead like to know the total Calories carried by the top three Elves carrying the most Calories.
That way, even if one of those Elves runs out of snacks, they still have two backups.

In the example above, the top three Elves are the fourth Elf (with 24000 Calories), then the third Elf (with 11000 Calories),
then the fifth Elf (with 10000 Calories). The sum of the Calories carried by these three elves is 45000.

Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
*/
func OneB() int {
	inputBytes, err := os.ReadFile("./one/input.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	input := string(inputBytes)
	// replace all CRLF's with LF's:
	// https://developer.mozilla.org/en-US/docs/Glossary/CRLF
	input = strings.ReplaceAll(input, "\r\n", "\n")

	inventories := strings.Split(input, "\n\n")

	textBlockSums := make([]int, len(inventories))

	// for each invetory/text block, we calculate it's total
	// and put that total in the textBlockSums array
	for i, inv := range inventories {
		entries := strings.Split(inv, "\n")

		var textBlockSum int
		for _, entry := range entries {

			entryValue, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatalf("Atoi error: %v", err)
			}

			textBlockSum += entryValue
		}

		textBlockSums[i] = textBlockSum

	}

	// we sort the array in place (largest to smallest)
	sort.Sort(sort.Reverse(sort.IntSlice(textBlockSums)))

	// return the sum of the first 3, or the 3 largest.
	// since we have the input data, we know we can safely access all 3 indeces.
	return textBlockSums[0] + textBlockSums[1] + textBlockSums[2]

}
