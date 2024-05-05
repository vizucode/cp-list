package reversedstring

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name     string
		args     args
		wantResp string
	}{
		{
			name: "Test 1",
			args: args{
				word: "hello",
			},
			wantResp: "olleh",
		},
		{
			name: "Test 2",
			args: args{
				word: "World",
			},
			wantResp: "dlroW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := ReverseString(tt.args.word); gotResp != tt.wantResp {
				t.Errorf("ReverseString() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
