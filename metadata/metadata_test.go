package metadata

import (
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func TestNew(t *testing.T) {
	type args struct {
		rows int
		cols int
	}
	tests := []struct {
		name string
		args args
		want metadata
	}{
		{
			name: "Zeros",
			args: args{0, 0},
			want: metadata{},
		},
		{
			name: "Zero Rows",
			args: args{0, 12323},
			want: metadata{},
		},
		{
			name: "Zero Cols",
			args: args{355756, 0},
			want: metadata{},
		},
		{
			name: "Valid",
			args: args{5, 4},
			want: metadata{
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metadata_Save(t *testing.T) {
	iss := is.New(t)
	tests := []struct {
		name     string
		m        metadata
		expected metadata
		err      error
		args     Slice
	}{
		{"Valid",
			metadata{
				{1, 2, 2, 3},
				{1, 2, 2, 3},
				{1, 2, 2, 3},
			},
			metadata{
				{1, 2, 2, 3},
				{1, 2, 0, 0},
				{1, 2, 0, 0},
			},
			nil,
			Slice{1, 2, 2, 3},
		},
		{"Zero size",
			metadata{},
			metadata{},
			ErrOutOfIndex,
			Slice{2, 3, 2, 3},
		},
		{"Single column",
			metadata{
				{1, 2, 2, 3},
				{1, 2, 2, 3},
				{1, 2, 2, 3},
			},
			metadata{
				{1, 2, 2, 3},
				{1, 2, 2, 1},
				{1, 2, 2, 1},
			},
			nil,
			Slice{1, 3, 2, 3},
		},
		{"Single row",
			metadata{
				{1, 2, 2, 3},
				{1, 2, 2, 3},
				{1, 2, 2, 3},
			},
			metadata{
				{1, 2, 2, 3},
				{1, 2, 2, 3},
				{1, 2, 2, 2},
			},
			nil,
			Slice{2, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Save(tt.args)
			iss.Equal(tt.m, tt.expected)
		})
	}
}
