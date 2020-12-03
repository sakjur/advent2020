package advent2020

import "fmt"

type XY struct {
	X int
	Y int
}

type Toboggan struct {
	Pos XY
}

type TreeMap struct {
	width  int
	height int
	m      map[XY]bool
}

func (m *TreeMap) AddLine(line string) error {
	if m.width == 0 {
		m.width = len(line)
	}

	if len(line) != m.width {
		return fmt.Errorf("expected width %d, got %d", m.width, len(line))
	}

	for i, c := range line {
		if c == '#' {
			m.m[XY{
				X: i,
				Y: m.height,
			}] = true
		}
	}

	m.height++
	return nil
}

func (m TreeMap) IsTree(p XY) bool {
	p.X = p.X % m.width

	return m.m[p]
}

func Travel(m TreeMap, slope XY) int {
	count := 0
	pos := XY{}
	for ; pos.Y <= m.height; pos.Y += slope.Y {
		if m.IsTree(pos) {
			count++
		}
		pos.X += slope.X
	}
	return count
}

func CreateMap(lines <-chan string) TreeMap {
	m := &TreeMap{
		m: map[XY]bool{},
	}
	for line := range lines {
		if err := m.AddLine(line); err != nil {
			panic(err)
		}
	}

	return *m
}
