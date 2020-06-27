package queens_attack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueensAttack(t *testing.T) {
	assert.Equal(
		t,
		int32(16),
		queensAttack(5, 0, 3, 3, [][]int32{}),
	)
	assert.Equal(
		t,
		int32(2),
		queensAttack(5, 0, 3, 3, [][]int32{
			{4, 3}, {5, 3}, //horiz right
			{2, 3}, {1, 3}, //horiz left
			{3, 4}, {3, 5}, //vert up
			{3, 2}, {3, 1}, //vert down
			{4, 4}, {5, 5}, //diag right up
			{2, 2}, {1, 1}, //diag right down
			{2, 4}, {1, 5}, //diag left up
			//{4, 2},{5, 1},//diag left down
		}),
	)
	assert.Equal(
		t,
		int32(11),
		queensAttack(5, 0, 4, 2, [][]int32{{3, 3}}),
	)
	assert.Equal(
		t,
		int32(10),
		queensAttack(5, 0, 4, 3, [][]int32{
			{5, 5},
			{4, 2},
			{2, 3},
		}),
	)
	assert.Equal(
		t,
		int32(6),
		queensAttack(5, 0, 3, 3, [][]int32{
			{4, 4},
			{5, 5},
			{3, 4},
			{3, 2},
			{4, 3},
			{2, 3}},
		),
	)
}
