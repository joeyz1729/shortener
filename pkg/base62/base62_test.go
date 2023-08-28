package base62

import "testing"

func TestInt2String(t *testing.T) {
	type args struct {
		seq uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "0",
			args: args{seq:  0},
			want: "0",
		},
		{
			name: "1",
			args: args{seq:  1},
			want: "1",
		},
		{
			name: "62",
			args: args{seq:  62},
			want: "10",
		},
		{
			name: "6347",
			args: args{seq:  6347},
			want: "1En",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2String(tt.args.seq); got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}
