package flyweight

import (
	"reflect"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	board1 := NewChessBoard()
	board1.Move(1, 1, 2) // 将棋子 1 移动到 (1, 2) 的位置
	board2 := NewChessBoard()
	board2.Move(2, 2, 3) // 将棋子 2 移动到 (2, 3) 的位置

	// 两个棋盘「炮」的享元部分是否相等
	if !reflect.DeepEqual(board1.chessPieces[2].Unit, board2.chessPieces[2].Unit) {
		t.Error("not equal 1")
	}
	// 两个棋盘「车」的享元部分是否相等
	if !reflect.DeepEqual(board1.chessPieces[1].Unit, board2.chessPieces[1].Unit) {
		t.Error("not equal 2")
	}
}
