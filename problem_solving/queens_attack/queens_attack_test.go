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
			name: "",
			args: args{
				posX:  4,
				posY:  4,
				sizeX: 8,
				sizeY: 8,
			},
			want: 27,
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
