package six

func numberOfCharsProcessedBeforeDistinctChars(input string, distinctChars int) int {
	for i := 0; i < len(input)-distinctChars; i++ {
		chars := make(map[byte]bool)

		for j := 0; j < distinctChars; j++ {
			chars[input[i+j]] = true
		}

		if len(chars) < distinctChars {
			continue
		}

		return len(input[:i+distinctChars])
	}

	return -1
}
