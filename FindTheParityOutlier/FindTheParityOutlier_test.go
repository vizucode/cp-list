package kata

import (
	"testing"
)

func TestFindOutlier(t *testing.T) {
	type args struct {
		integers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				integers: []int{2, 4, 0, 100, 4, 11, 2602, 36},
			},
			want: 11,
		},
		{
			name: "test case 2",
			args: args{
				integers: []int{11, 1, 3, 7, 9, 12, 7, 9, 5},
			},
			want: 12,
		},
		{
			name: "test case 3",
			args: args{
				integers: []int{1, 11, 13, 17, 18, 12, 7, 9, 5},
			},
			want: 18,
		},
		{
			name: "test case 4",
			args: args{
				integers: []int{-2, 3, 1, 5, -5, 7},
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOutlier(tt.args.integers); got != tt.want {
				t.Errorf("FindOutlier() = %v, want %v", got, tt.want)
			}
		})
	}
}
