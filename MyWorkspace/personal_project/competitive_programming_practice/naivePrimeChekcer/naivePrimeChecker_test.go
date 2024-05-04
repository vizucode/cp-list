package naiveprimechekcer

import "testing"

func TestNaivePrimeChecker(t *testing.T) {
	type args struct {
		N uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				N: 7123712939882223322,
			},
			want: false,
		},
		{
			name: "test case 2",
			args: args{
				N: 14889444574204132552,
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				N: 14889444523123325522,
			},
			want: false,
		},
		{
			name: "test case 4",
			args: args{
				N: 1000,
			},
			want: false,
		},
		{
			name: "test case 5",
			args: args{
				N: 7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NaivePrimeChecker(tt.args.N); got != tt.want {
				t.Errorf("NaivePrimeChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
