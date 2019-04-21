package calc

import "testing"

func TestMult(t *testing.T) {
	type args struct {
		args []int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{
			name:       "Able to execute happy path",
			args:       args{[]int64{1, 2, 3}},
			wantResult: 6,
		},
		{
			name:       "Handle with negative values",
			args:       args{[]int64{1, -2, 3}},
			wantResult: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Mult(tt.args.args); gotResult != tt.wantResult {
				t.Errorf("Mult() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
