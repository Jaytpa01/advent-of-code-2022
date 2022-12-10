package nine

import (
	"fmt"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func B() int {

	input := utils.NewFileScanner("./9/input.txt")
	defer input.Close()
	// input := utils.NewFileScanner("./9/sample.txt")

	rope := make([]*point, 10)
	for i := range rope {
		rope[i] = &point{1, 1}
	}
	head := rope[0]
	tail := rope[9]

	tailVisited := map[string]bool{}

	for input.Scan() {
		move := input.Text()

		var direction string
		var steps int

		fmt.Sscanf(move, "%s %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			head.move(direction)

			for i := 1; i < len(rope); i++ {
				if !rope[i].isTouching(rope[i-1]) {
					rope[i].follow2(rope[i-1])
				}
			}

			tailVisited[tail.String()] = true

		}

	}

	return len(tailVisited)
}

func (a *point) follow2(b *point) {

	x := b.x - a.x
	y := b.y - a.y

	if x > 0 {
		a.x++
	} else if x < 0 {
		a.x--
	}

	if y > 0 {
		a.y++
	} else if y < 0 {
		a.y--
	}

}
