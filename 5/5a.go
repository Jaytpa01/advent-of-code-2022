package five

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

/*
--- Day 5: Supply Stacks ---
The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3
In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3
Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3
Finally, one crate is moved from stack 1 to stack 2:

        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3
The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each stack?
*/

// my input stacks
// [N] [G]                     [Q]
// [H] [B]         [B] [R]     [H]
// [S] [N]     [Q] [M] [T]     [Z]
// [J] [T]     [R] [V] [H]     [R] [S]
// [F] [Q]     [W] [T] [V] [J] [V] [M]
// [W] [P] [V] [S] [F] [B] [Q] [J] [H]
// [T] [R] [Q] [B] [D] [D] [B] [N] [N]
// [D] [H] [L] [N] [N] [M] [D] [D] [B]
//  1   2   3   4   5   6   7   8   9

// sample stack
//
//	[D]
//
// [N] [C]
// [Z] [M] [P]
//
//	1   2   3
func A() string {

	reg := regexp.MustCompile(`(\d{1}) from (\d{1}) to (\d{1})`)

	stack1 := &stack{"D", "T", "W", "F", "J", "S", "H", "N"}
	stack2 := &stack{"H", "R", "P", "Q", "T", "N", "B", "G"}
	stack3 := &stack{"L", "Q", "V"}
	stack4 := &stack{"N", "B", "S", "W", "R", "Q"}
	stack5 := &stack{"N", "D", "F", "T", "V", "M", "B"}
	stack6 := &stack{"M", "D", "B", "V", "H", "T", "R"}
	stack7 := &stack{"D", "B", "Q", "J"}
	stack8 := &stack{"D", "N", "J", "V", "R", "Z", "H", "Q"}
	stack9 := &stack{"B", "N", "H", "M", "S"}
	stacks := []*stack{stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9}
	scanner := utils.NewFileScanner("./5/input.txt")

	// stack1 := &stack{"Z", "N"}
	// stack2 := &stack{"M", "C", "D"}
	// stack3 := &stack{"P"}
	// stacks := []*stack{stack1, stack2, stack3}
	// scanner := utils.NewFileScanner("./5/sample.txt")

	defer scanner.Close()
	iter := 1

	for scanner.Scan() {
		procedure := scanner.Text()

		instructions := reg.FindStringSubmatch(procedure)
		quantity, _ := strconv.Atoi(instructions[1])
		from, _ := strconv.Atoi(instructions[2])
		to, _ := strconv.Atoi(instructions[3])

		for i := 0; i < quantity; i++ {
			crate := stacks[from-1].Pop()

			if crate != "" {
				stacks[to-1].Push(crate)
			} else {
				fmt.Println(procedure)
				fmt.Println(iter)

			}
		}

		iter++
	}

	// for i := range stacks {
	// 	fmt.Println(stacks[i])
	// }

	fmt.Println()
	output := ""
	for _, stk := range stacks {
		output += stk.Pop()

	}

	return output
}

type stack []string

func (s *stack) Push(in string) {
	*s = append(*s, in)
}

func (s *stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}

	// find the last index and get the element there
	lastIdx := len(*s) - 1
	elem := (*s)[lastIdx]

	// remove last item from stack
	*s = (*s)[:lastIdx]

	// return item
	return elem
}
