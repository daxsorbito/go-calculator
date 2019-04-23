package calc

import "testing"

func TestDiv(t *testing.T) {
	type args struct {
		args []int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult float32
	}{
		{
			name:       "Able to execute properly",
			args:       args{[]int64{10, 5}},
			wantResult: 2,
		},
		{
			name:       "Handle with negative values",
			args:       args{[]int64{10, -5}},
			wantResult: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Div(tt.args.args); gotResult != tt.wantResult {
				t.Errorf("Div() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
