package metadata

import "errors"

// metadata is type to represent current processing state [rows][cols]int
// -1 is unused element
// 0-N - ids of slices
type metadata [][]int

type Slice struct {
	R1 int
	C1 int
	R2 int
	C2 int
}

var ErrOutOfIndex = errors.New("not valid index")

var currentID int

func New(rows, cols int) metadata {
	if rows == 0 || cols == 0 {
		return metadata{}
	}
	data := make([][]int, rows)

	for r := 0; r < rows; r++ {
		data[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			data[r][c] = -1
		}
	}
	return data
}

func (m metadata) Save(s Slice) error {
	if s.R1 > s.R2 {
		s.R1, s.R2 = s.R2, s.R1
	}
	if s.C1 > s.C2 {
		s.C1, s.C2 = s.C2, s.C1
	}

	if s.R2 > len(m) || s.C2 > len(m[0]) {
		return ErrOutOfIndex
	}

	for r := s.R1; r <= s.R2; r++ {
		for c := s.C1; c <= s.C2; c++ {
			m[r][c] = currentID
		}
	}

	currentID++
	return nil
}

func (m metadata) Get(row, col int) Slice {
	var slice Slice
	found := [][2]int{{row, col}}
	for _, cords := range found {
		found = append(found, m.findSame(cords[0], cords[1])...)

		// To find coordinates of high angles
		switch {
		case slice.R1 > cords[0]:
			slice.R1 = cords[0]
		case slice.R2 < cords[0]:
			slice.R2 = cords[0]
		case slice.C1 > cords[1]:
			slice.C1 = cords[1]
		case slice.C2 < cords[1]:
			slice.C2 = cords[1]
		}
	}
	return slice
}

func (m metadata) findSame(row, col int) [][2]int {
	var foundElements [][2]int
	id := m[row][col]
	toCheck := [4][2]int{{row, col - 1}, {row, col + 1}, {row - 1, col}, {row + 1, col}}

	for _, cords := range toCheck {
		if m[cords[0]][cords[1]] == id {
			foundElements = append(foundElements, [2]int{cords[0], cords[1]})
		}
	}
	return foundElements
}
