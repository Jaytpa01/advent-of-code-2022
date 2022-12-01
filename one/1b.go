package one

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func OneB() {
	inputBytes, err := os.ReadFile("./one/input.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	input := string(inputBytes)
	input = strings.ReplaceAll(input, "\r\n", "\n")

	inventories := strings.Split(input, "\n\n")

	sums := make([]int, len(inventories))

	for i, inv := range inventories {
		entries := strings.Split(inv, "\n")

		var sum int
		for _, entry := range entries {
			entInt, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatalf("Atoi error: %v", err)
			}

			sum += entInt
		}

		sums[i] = sum

	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	total := sums[0] + sums[1] + sums[2]
	fmt.Println(total)

}
