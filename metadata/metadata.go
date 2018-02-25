package metadata

import (
	"errors"
	"sync"
)

// metadata is type to represent current processing state [rows][cols]int
// -1 is unused element
// 0-N - ids of slices
type metadata struct {
	currentID int
	mu        *sync.RWMutex

	Index  [][]int
	Slices []Slice
}

type Slice struct {
	ID int
	R1 int
	C1 int
	R2 int
	C2 int
}

const Empty = -1

var ErrOutOfIndex = errors.New("not valid index")

func New(rows, cols int) metadata {
	if rows == 0 || cols == 0 {
		return metadata{}
	}
	data := make([][]int, rows)

	for r := 0; r < rows; r++ {
		data[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			data[r][c] = Empty
		}
	}
	return metadata{Index: data, mu: &sync.RWMutex{}}
}

func (m *metadata) Save(s Slice) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if s.R1 > s.R2 {
		s.R1, s.R2 = s.R2, s.R1
	}

	if s.C1 > s.C2 {
		s.C1, s.C2 = s.C2, s.C1
	}

	if s.R2 > len(m.Index) || s.C2 > len(m.Index[0]) {
		return ErrOutOfIndex
	}

	for r := s.R1; r <= s.R2; r++ {
		for c := s.C1; c <= s.C2; c++ {
			m.Index[r][c] = m.currentID
		}
	}
	s.ID = m.currentID
	m.Slices = append(m.Slices, s)

	m.currentID++
	return nil
}

func (m metadata) Get(row, col int) Slice {
	m.mu.RLock()
	defer m.mu.RUnlock()

	id := m.Index[row][col]
	if id == Empty {
		return Slice{Empty, row, col, row, col}
	}

	return m.Slices[id]
}

func (m metadata) IsEmpty(s Slice) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if s.R1 > s.R2 {
		s.R1, s.R2 = s.R2, s.R1
	}

	if s.C1 > s.C2 {
		s.C1, s.C2 = s.C2, s.C1
	}

	if s.R2 > len(m.Index) || s.C2 > len(m.Index[0]) {
		return false, ErrOutOfIndex
	}

	for r := s.R1; r <= s.R2; r++ {
		for c := s.C1; c <= s.C2; c++ {
			if m.Index[r][c] != Empty {
				return false, nil
			}
		}
	}
	return true, nil
}

func (m *metadata) correct(maxCells int) {
	for _, v := range m.Slices {
		// TODO update slice if substitution was made
		// TODO iterate over slices till no more changes can be made
		gogo(v, m.Index, maxCells)
	}
}
