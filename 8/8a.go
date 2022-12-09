package eight

import (
	"strconv"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func A() int {

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

	visible := (len(treeGrid) * 2) + ((len(treeGrid[0]) - 2) * 2)

	for i := 1; i < len(treeGrid)-1; i++ {
		for j := 1; j < len(treeGrid[i])-1; j++ {

			currentTree := treeGrid[i][j]

			rightVisible := true
			leftVisible := true
			upVisible := true
			downVisible := true

			for hor := 0; hor < len(treeGrid[i]); hor++ {
				if hor == j {
					continue
				}

				if treeGrid[i][hor] >= currentTree {
					if hor < j {
						leftVisible = false
					} else {
						rightVisible = false
					}
				}
			}

			for vert := 0; vert < len(treeGrid); vert++ {
				if vert == i {
					continue
				}

				if treeGrid[vert][j] >= currentTree {
					if vert < i {
						upVisible = false
					} else {
						downVisible = false
					}
				}
			}

			if leftVisible || rightVisible || upVisible || downVisible {
				visible++
			}

		}
	}

	return visible
}
