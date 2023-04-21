package flyweight

// 棋子是不可变的
var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "車",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
	// ... 其他棋子
}

// 将不可变的棋子享元
type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

// 享元工厂，用来生产棋子享元
func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

// 棋子，棋子分为可以共享的内部状态和不可共享的外部状态
// 这里可以共享的部分是棋子，不可共享的是棋子位置
type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

// 棋局，是映射
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

// 初始化棋盘
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

// 移动棋子
func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}
