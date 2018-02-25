package metadata

import (
	"reflect"
	"testing"

	"sync"

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
				Index: [][]int{
					{-1, -1, -1, -1},
					{-1, -1, -1, -1},
					{-1, -1, -1, -1},
					{-1, -1, -1, -1},
					{-1, -1, -1, -1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got.Index, tt.want.Index) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}

			if got := New(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got.Slices, tt.want.Slices) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metadata_Save(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		name     string
		m        metadata
		expected metadata
		err      error
		args     Slice
	}{
		{"Valid",
			metadata{
				mu:        &sync.RWMutex{},
				currentID: 0,
				Index: [][]int{
					{1, 2, 2, 3},
					{1, 2, 2, 3},
					{1, 2, 2, 3},
				},
			},
			metadata{
				currentID: 1,
				Index: [][]int{
					{1, 2, 2, 3},
					{1, 2, 0, 0},
					{1, 2, 0, 0},
				},
			},
			nil,
			Slice{0, 1, 2, 2, 3},
		},
		{"Zero size",
			metadata{mu: &sync.RWMutex{}},
			metadata{},
			ErrOutOfIndex,
			Slice{0, 2, 3, 2, 3},
		},
		{"Single column",
			metadata{
				mu:        &sync.RWMutex{},
				currentID: 1,
				Index: [][]int{
					{1, 2, 2, 3},
					{1, 2, 2, 3},
					{1, 2, 2, 3},
				},
			},
			metadata{
				currentID: 2,
				Index: [][]int{
					{1, 2, 2, 3},
					{1, 2, 2, 1},
					{1, 2, 2, 1},
				},
			},
			nil,
			Slice{0, 1, 3, 2, 3},
		},
		{"Single row",
			metadata{
				mu:        &sync.RWMutex{},
				currentID: 2,
				Index: [][]int{
					{1, 2, 2, 3},
					{1, 2, 2, 3},
					{1, 2, 2, 3},
				},
			},
			metadata{
				currentID: 3,
				Index: [][]int{
					{1, 2, 2, 3},
					{1, 2, 2, 3},
					{1, 2, 2, 2},
				},
			},
			nil,
			Slice{0, 2, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Save(tt.args)
			if !reflect.DeepEqual(tt.m.Index, tt.expected.Index) {
				t.Errorf("Expected %+v, got %+v", tt.expected.Index, tt.m.Index)
			}
			is.Equal(tt.m.currentID, tt.expected.currentID)
		})
	}
}
