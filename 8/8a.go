package eight

import (
	"fmt"
	"strconv"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func A() int {

	input := utils.NewFileScanner("./8/sample.txt")
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

	// for _, trees := range treeGrid {
	// 	fmt.Println(trees)
	// }
	visibleInterior := 0

	for i := 1; i < len(treeGrid)-1; i++ {
		for j := 1; j < len(treeGrid[i])-1; j++ {

			currentTree := treeGrid[i][j]

			fmt.Println(currentTree, i, j)

			rightVisible := true
			leftVisible := true
			upVisible := true
			downVisible := true

			// check right
			for right := 1; right < len(treeGrid[i])-1; right++ {
				if treeGrid[i][j+right] >= currentTree {
					rightVisible = false
					break
				}
			}

			// check left
			for left := j; left >= 0; left-- {
				if treeGrid[i][j-left] >= currentTree {
					leftVisible = false
					break
				}
			}

			// check up
			for up := i; up >= 0; up-- {
				if treeGrid[i-up][j] >= currentTree {
					upVisible = false
					break
				}
			}

			// check down
			for down := 1; down <= len(treeGrid)-1; down++ {
				if treeGrid[i+down][j] >= currentTree {
					downVisible = false
					break
				}
			}

			if leftVisible || rightVisible || upVisible || downVisible {
				visibleInterior++
			}

		}
	}

	fmt.Println(visibleInterior)

	return 0
}
