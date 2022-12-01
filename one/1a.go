package one

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func OneA() {
	inputBytes, err := os.ReadFile("./one/input.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	input := string(inputBytes)
	input = strings.ReplaceAll(input, "\r\n", "\n")

	inventories := strings.Split(input, "\n\n")

	var largestSum int

	for _, inv := range inventories {
		entries := strings.Split(inv, "\n")

		var sum int
		for _, entry := range entries {
			entInt, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatalf("Atoi error: %v", err)
			}

			sum += entInt
		}

		if sum > largestSum {
			largestSum = sum
		}

	}

	fmt.Println(largestSum)

}
