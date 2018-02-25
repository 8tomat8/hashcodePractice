package input

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Data will contins input information about pizza
type Data struct {
	R, C, L, H int
	P          [][]bool // pizza representation, tomat - true, mashrooms - false
}

// Fill will read pizza input information to the Data struct
func (d *Data) Fill(src io.Reader) error {
	scanner := bufio.NewScanner(src)

	if ok := scanner.Scan(); !ok {
		err := scanner.Err()
		return errors.Wrap(err, "reading first line from input source of pizza data")
	}

	n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &d.R, &d.C, &d.L, &d.H)
	if err != nil || n != 4 {
		return errors.Wrapf(err, "first line missed some important valuest, should contain 4 separate number (%d was readed)", n)
	}

	d.P = make([][]bool, d.R)
	rowSlice := make([]byte, d.C)
	i := 0 // row index
	for scanner.Scan() {
		d.P[i] = make([]bool, d.C)
		rowSlice = scanner.Bytes()
		for j, ch := range rowSlice {
			if ch == 'T' {
				d.P[i][j] = true
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "scanning input pizza data")
	}

	return nil
}
