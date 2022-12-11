package eleven

type monkey struct {
	id    int
	items itemQueue
}

type itemQueue []int

func (i *itemQueue) enqueue(item int) {
	*i = append(*i, item)
}

func (i *itemQueue) dequeue() int {
	item := (*i)[0]
	*i = (*i)[1:]

	return item
}
