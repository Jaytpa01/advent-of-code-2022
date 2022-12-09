package eight

import (
	"strconv"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func B() int {

	input := utils.NewFileScanner("./8/input.txt")
	defer input.Close()

	treeGrid := [][]int{}

	for input.Scan() {
		line := input.Text()

		treeLine := []int{}
		for _, char := range line {
			i, _ := strconv.Atoi(string(char))
			treeLine = append(treeLine, i)
		}

		treeGrid = append(treeGrid, treeLine)
	}

	largestScenicScore := 0

	for i := 0; i < len(treeGrid); i++ {
		for j := 0; j < len(treeGrid[i]); j++ {

			currentTree := treeGrid[i][j]

			// search left
			leftCount := 0
			for left := 1; left <= len(treeGrid[i][:j]); left++ {
				lookedAt := treeGrid[i][j-left]
				if lookedAt < currentTree {
					leftCount++
					continue
				}

				leftCount++
				break
			}

			rightCount := 0
			for right := 1; right < len(treeGrid[i][j:]); right++ {
				lookedAt := treeGrid[i][j+right]
				if lookedAt < currentTree {
					rightCount++
					continue
				}

				rightCount++
				break
			}

			upCount := 0
			for up := 1; up <= len(treeGrid[:i]); up++ {
				lookedAt := treeGrid[i-up][j]
				if lookedAt < currentTree {
					upCount++
					continue
				}

				upCount++
				break
			}

			downCount := 0
			for down := 1; down < len(treeGrid[i:]); down++ {
				lookedAt := treeGrid[i+down][j]
				if lookedAt < currentTree {
					downCount++
					continue
				}

				downCount++
				break
			}

			score := rightCount * leftCount * upCount * downCount
			if score > largestScenicScore {
				largestScenicScore = score
			}

		}
	}

	return largestScenicScore
}
