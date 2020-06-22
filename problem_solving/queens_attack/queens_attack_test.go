package queens_attack

import "testing"

func TestQueensAttack(t *testing.T) {
	type args struct {
		posX      int32
		posY      int32
		sizeX     int32
		sizeY     int32
		obstacles [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Default test case #1",
			args: args{
				posX:  8,
				posY:  1,
				sizeX: 4,
				sizeY: 4,
				obstacles: [][]int32{
					{3, 5},
				},
			},
			want: 24,
		},
		{
			name: "Default test case",
			args: args{
				posX:  4,
				posY:  0,
				sizeX: 4,
				sizeY: 0,
			},
			want: 9,
		},
		{
			name: "Default test case #2",
			args: args{
				posX:  5,
				posY:  3,
				sizeX: 4,
				sizeY: 3,
				obstacles: [][]int32{
					{5, 5},
					{4, 2},
					{2, 3},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queensAttack(tt.args.posX, tt.args.posY, tt.args.sizeX, tt.args.sizeY, tt.args.obstacles); got != tt.want {
				t.Errorf("queensAttack() = %v, want %v", got, tt.want)
			}
		})
	}
}
