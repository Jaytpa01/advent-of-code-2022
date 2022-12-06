package five

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

func (s *stack) PopMultiple(qty int) stack {
	// fmt.Println(s)

	if len(*s) == 0 {
		return stack{}
	}

	// find the last items and get the element there
	idx := len(*s)
	elems := (*s)[idx-qty:]

	// remove last items from stack
	*s = (*s)[:idx-qty]

	// fmt.Println(s)

	// return item
	return elems
}

func (s *stack) PushMultiple(in stack) {
	*s = append(*s, in...)
}

func (s *stack) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
