package binarycount

import "testing"

func TestBinaryCount(t *testing.T) {
	type args struct {
		N uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Test 1",
			args: args{
				N: 1234,
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				N: 20,
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				N: 5000,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryCount(tt.args.N); got != tt.want {
				t.Errorf("BinaryCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
