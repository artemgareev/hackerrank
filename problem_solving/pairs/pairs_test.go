package pairs

import "testing"

func Test_pairs(t *testing.T) {
	type args struct {
		k   int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"test_case_0",
			args{2, []int32{1, 3, 5, 8, 6, 4, 2}},
			5,
		},
		{
			"test_case_1",
			args{1, []int32{363374326, 364147530, 61825163, 1073065718, 1281246024, 1399469912, 428047635, 491595254, 879792181, 1069262793}},
			0,
		},
		{
			"test_case_2",
			args{2, []int32{1, 5, 3, 4, 2}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairs(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("pairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
