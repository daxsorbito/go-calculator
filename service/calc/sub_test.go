package calc

import "testing"

func TestSub(t *testing.T) {
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
			args:       args{[]int64{1, 2, 3}},
			wantResult: -6,
		},
		{
			name:       "Handle with negative values",
			args:       args{[]int64{1, -2, 3}},
			wantResult: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Sub(tt.args.args); gotResult != tt.wantResult {
				t.Errorf("Sub() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
