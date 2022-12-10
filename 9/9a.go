package nine

import (
	"fmt"
	"math"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func A() int {

	input := utils.NewFileScanner("./9/input.txt")
	// input := utils.NewFileScanner("./9/sample.txt")
	defer input.Close()

	head := &point{1, 1}
	tail := &point{1, 1}

	tailVisited := map[string]bool{}

	for input.Scan() {
		move := input.Text()

		var direction string
		var steps int

		fmt.Sscanf(move, "%s %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			head.move(direction)

			if !tail.isTouching(head) {
				tail.follow(head, direction)
			}

			tailVisited[tail.String()] = true

		}

	}

	return len(tailVisited)
}

type point struct {
	x int
	y int
}

func (p *point) move(dir string) {
	switch dir {
	case "U":
		p.y++
	case "D":
		p.y--
	case "L":
		p.x--
	case "R":
		p.x++
	}
}

func (a *point) follow(b *point, bMoveDirection string) {

	// if b moves horizontally / is on same row
	if a.y == b.y {
		switch bMoveDirection {
		case "R":
			a.x++
		case "L":
			a.x--
		}

		return
	}

	// if b moved vertically / is on same column
	if a.x == b.x {
		switch bMoveDirection {
		case "U":
			a.y++
		case "D":
			a.y--
		}

		return
	}

	// not on same column or row, need to move horizontally
	switch bMoveDirection {
	case "U":
		if b.x > a.x {
			a.x++
		} else {
			a.x--
		}

		a.y++
	case "D":
		if b.x > a.x {
			a.x++
		} else {
			a.x--
		}

		a.y--
	case "R":
		if b.y > a.y {
			a.y++
		} else {
			a.y--
		}

		a.x++
	case "L":
		if b.y > a.y {
			a.y++
		} else {
			a.y--
		}
		a.x--
	}

}

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (a *point) isTouching(b *point) bool {
	x := math.Abs(float64(a.x - b.x))
	y := math.Abs(float64(a.y - b.y))

	return x >= 0 && x <= 1 && y >= 0 && y <= 1
}
