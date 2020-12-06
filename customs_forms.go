package advent2020

type CustomsGroup struct {
	groupSize int
	answers   map[rune]int
}

func (g CustomsGroup) Any() int {
	return len(g.answers)
}

func (g CustomsGroup) All() int {
	count := 0
	for _, answers := range g.answers {
		if answers == g.groupSize {
			count++
		}
	}
	return count
}

func CustomForms(c <-chan string) []CustomsGroup {
	result := []CustomsGroup{}
	current := map[rune]int{}
	groupSize := 0
	for line := range c {
		if line == "" {
			result = append(result, CustomsGroup{
				groupSize: groupSize,
				answers:   current,
			})
			groupSize = 0
			current = map[rune]int{}
			continue
		}
		groupSize++

		for _, c := range line {
			if _, exists := current[c]; exists {
				current[c]++
			} else {
				current[c] = 1
			}
		}
	}
	result = append(result, CustomsGroup{
		groupSize: groupSize,
		answers:   current,
	})
	return result
}
