package countoddnumbern

import "testing"

func TestCountOddNumberN(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name     string
		args     args
		wantResp int
	}{
		{
			name: "Test 1",
			args: args{
				N: 5,
			},
			wantResp: 2,
		},
		{
			name: "test 2",
			args: args{
				N: 10,
			},
			wantResp: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := CountOddNumberN(tt.args.N); gotResp != tt.wantResp {
				t.Errorf("CountOddNumberN() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
