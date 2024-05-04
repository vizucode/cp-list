package switchlamp

import "testing"

func Test_switchLampBruteForce(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				N: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := switchLampBruteForce(tt.args.N); got != tt.want {
				t.Errorf("switchLampBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
