package input

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestFill(t *testing.T) {
	ist := is.New(t)
	in := []byte(`3 5 1 6
TTTTT
TMMMT
TTTTT
`)

	d := Data{}
	err := d.Fill(bytes.NewReader(in))
	ist.NoErr(err) // genneral cases at input source
}
