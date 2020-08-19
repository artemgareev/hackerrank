package common_child

import "testing"

func Test_commonChild(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"test_case_0",
			args{"HARRY", "SALLY"},
			int32(2),
		},
		{
			"test_case_1",
			args{"SHINCHAN", "NOHARAAA"},
			int32(3),
		},
		{
			"test_case_2",
			args{"AA", "BB"},
			int32(0),
		},
		{
			"test_case_3",
			args{"ABCDEF", "FBDAMN"},
			int32(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChild(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("commonChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
