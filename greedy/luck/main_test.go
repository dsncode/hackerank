package main

import "testing"

func TestLuck(t *testing.T) {

	// [[5, 1], [2, 1], [1, 1], [8, 1], [10, 0], [5, 0]]
	contests := [][]int32{
		{5, 1},
		{2, 1},
		{1, 1},
		{8, 1},
		{10, 0},
		{5, 0},
	}
	var k int32 = 3
	luck := luckBalance(k, contests)

	if luck != 29 {
		t.Fatal("wrong output")
	}

}
