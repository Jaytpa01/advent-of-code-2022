package eleven

import "strconv"

type monkey struct {
	items       itemQueue
	calculation calculation
	divisibleBy int

	ifTrue  int
	ifFalse int

	inspections int
}

func (m *monkey) performOperation(old int) int {

	switch m.calculation.operation {
	case "*":
		if m.calculation.val == "old" {
			return old * old
		}

		val, _ := strconv.Atoi(m.calculation.val)
		return old * val
	case "+":
		if m.calculation.val == "old" {
			return old + old
		}

		val, _ := strconv.Atoi(m.calculation.val)
		return old + val

	}

	return -1
}

type itemQueue []int

type calculation struct {
	operation string
	val       string
}

func (i *itemQueue) enqueue(item int) {
	*i = append(*i, item)
}

func (i *itemQueue) dequeue() int {
	item := (*i)[0]
	*i = (*i)[1:]

	return item
}
