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

func BenchmarkFill(b *testing.B) {
	in := []byte(`6 7 1 5
TMMMTTT
MMMMTMM
TTMTTMT
TMMTMMM
TTTTTTM
TTTTTTM
`)
	d := Data{}
	r := bytes.NewReader(in)

	for i := 0; i < b.N; i++ {
		d.Fill(r)
	}
}
