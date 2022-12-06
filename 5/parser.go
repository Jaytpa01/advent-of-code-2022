package five

import (
	"strconv"
	"strings"
)

// [N] [G]                     [Q]
// [H] [B]         [B] [R]     [H]
// [S] [N]     [Q] [M] [T]     [Z]
// [J] [T]     [R] [V] [H]     [R] [S]
// [F] [Q]     [W] [T] [V] [J] [V] [M]
// [W] [P] [V] [S] [F] [B] [Q] [J] [H]
// [T] [R] [Q] [B] [D] [D] [B] [N] [N]
// [D] [H] [L] [N] [N] [M] [D] [D] [B]
//  1   2   3   4   5   6   7   8   9

func parseCrates(input string) []*stack {
	rows := strings.Split(input, "\n")

	idxArr := strings.Split(rows[len(rows)-1], "")

	nums := strings.ReplaceAll(rows[len(rows)-1], " ", "")
	lastNum := string(nums[len(nums)-1])
	numOfCols, _ := strconv.Atoi(lastNum)

	stacks := make([]*stack, numOfCols)

	rows = rows[:len(rows)-1]
	for _, row := range rows {

		for i, char := range row {
			switch char {
			case '[', ']', ' ':
				continue
			}

			stackIdx, _ := strconv.Atoi(idxArr[i])
			stackIdx-- // decrement to go from 1 indexed to 0 indexed

			if stacks[stackIdx] == nil {
				stacks[stackIdx] = &stack{}
			}

			*stacks[stackIdx] = append(*stacks[stackIdx], string(char))

		}

	}

	for _, stk := range stacks {
		stk.Reverse()
	}

	return stacks

}
