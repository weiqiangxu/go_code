package p1

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 3, 9, 5, 7},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quickSort(tt.args.arr)
			t.Log(got)
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				numbers: []int{1, 8, 3, 9, 4, 5, 6},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bubbleSort(tt.args.numbers)
			t.Log(got)
		})
	}
}
