package main

import "errors"

type Solution struct {
	direction [][]int // 方向, 常量

	xLimit, yLimit int // 列的限制, 行的限制

	boards [][]byte // 棋牌
}

func newBoard(boards [][]byte) *Solution {
	return &Solution{
		boards:    boards,
		direction: [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}},
	}
}

func (b *Solution) inArea(x, y int) bool {
	return x >= 0 && y >= 0 && x < b.xLimit && y < b.yLimit
}

func exist(board [][]byte, word string) bool {
	yLimit := len(board)
	if yLimit == 0 {
		return false
	}

	xLimit := len(board[0])
	markd := make([][]bool, 0, 0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				moved = []location{{
					x: i,
					y: j,
				}}
			}

		}
	}

	if moved == nil {
		return false
	}

	for _, letter := range word {

	}

}


func main() {

}
