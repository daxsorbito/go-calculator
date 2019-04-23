package calc

import "testing"

func TestFac(t *testing.T) {
	type args struct {
		arg int64
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Able to execute properly",
			args: args{6},
			want: 720,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fac(tt.args.arg); got != tt.want {
				t.Errorf("Fac() = %v, want %v", got, tt.want)
			}
		})
	}
}
