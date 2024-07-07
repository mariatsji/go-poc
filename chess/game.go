package chess

import (
	"fmt"
	"go-poc/functional"
)

type Color int

const (
	White Color = iota
	Black
)

type Piece int

type ColoredPiece struct {
	Color Color
	Piece Piece
}

const (
	Pawn Piece = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func (p ColoredPiece) String() string {
	switch p.Piece {
	case Pawn:
		if p.Color == White {
			return "♙"
		} else {
			return "♟️"
		}
	case Knight:
		if p.Color == White {
			return "♘"
		} else {
			return "♞"
		}
	case Bishop:
		if p.Color == White {
			return "♗"
		} else {
			return "♝"
		}
	case Rook:
		if p.Color == White {
			return "♖"
		} else {
			return "♜"
		}
	case Queen:
		if p.Color == White {
			return "♕"
		} else {
			return "♛"
		}
	case King:
		if p.Color == White {
			return "♔"
		} else {
			return "♚"
		}
	default:
		return "?"
	}
}

type Board [12]uint64

func (b Board) String() string {
	retVal := ""
	for row := 7; row >= 0; row-- {
		for col := 0; col <= 7; col++ {
			pieceMaybe := SlowPieceAt(b, Square{row, col})
			if functional.Empty[ColoredPiece](pieceMaybe) {
				if (col+row)%2 == 0 {
					retVal += "▫"
				} else {
					retVal += "▪"
				}
			} else {
				retVal += functional.Get[ColoredPiece](pieceMaybe).String()
			}
		}
		retVal += "\n"
	}
	return retVal
}

func InsideBoard(square Square) bool {
	return square.Row > 0 && square.Row < 8 && square.Col > 0 && square.Col < 8
}

// set a piece at the board
func SetPiece(piece ColoredPiece, board Board, row, col int) Board {
	idx := int(piece.Piece)
	if piece.Color == Black {
		idx += 6
	}
	board[idx] = setBit(board[idx], row, col)
	return board
}

// remove a piece from the board
func ClearPiece(board Board, row, col int) Board {
	for pi := 0; pi < 12; pi++ {
		board[pi] = clearBit(board[pi], row, col)
	}
	return board
}

// make a move
// returns new Board - does not mutate incoming Board!
func Move(board Board, fromSquare Square, toSquare Square) Board {
	pieceMaybe := SlowPieceAt(board, fromSquare)
	functional.DoIf(pieceMaybe,
		(func(coloredPiece ColoredPiece) {
			clearedTo := ClearPiece(board, toSquare.Row, toSquare.Col)
			board = SetPiece(coloredPiece, clearedTo, toSquare.Row, toSquare.Col)
		}))
	return ClearPiece(board, fromSquare.Row, fromSquare.Col)
}

// setBit sets the bit at the given position on the chessboard.
func setBit(x uint64, row, col int) uint64 {
	x |= 1 << (row*8 + col)
	return x
}

// clearBit clears the bit at the given position on the chessboard.
func clearBit(x uint64, row, col int) uint64 {
	x &^= 1 << (row*8 + col)
	return x
}

// isSet checks if the bit at the given position on the chessboard is set.
func isSet(x uint64, row, col int) bool {
	return x&(1<<(row*8+col)) != 0
}

func SlowVacantAt(board Board, square Square) bool {
	return SlowPieceAt(board, square) == functional.None[ColoredPiece]()
}

func HasColorPieceAt(board Board, square Square, enemy Color) (ret bool) {
	functional.DoIf(
		SlowPieceAt(board, square),
		(func(p ColoredPiece) {
			if p.Color == enemy {
				ret = true
			}
		}),
	)
	return ret

}

func SlowPieceAt(board Board, square Square) functional.Maybe[ColoredPiece] {
	for pi := 0; pi < 12; pi++ {
		color := White
		if pi > 5 {
			color = Black
		}
		if isSet(board[pi], square.Row, square.Col) {
			piece := Piece(pi)
			if color == Black {
				piece = Piece(pi - 6)
			}
			return functional.Some(ColoredPiece{color, piece})
		}
	}
	return functional.None[ColoredPiece]()
}

// row and col from the int
func To(num uint64) (row, col int) {
	row = int(num) / 8 // Integer division gives the row.
	col = int(num) % 8 // Modulo operation gives the column.
	return
}

// int from a row and a col
func From(row, col int) uint64 {
	return uint64(row*8 + col)
}

type Square = struct {
	Row int
	Col int
}

func FindAll(board Board, piece Piece, color Color) []Square {
	idx := int(piece)
	if color == Black {
		idx += 6
	}
	n := board[idx]
	var positions []uint64
	for i := 0; i < 64; i++ {
		// Check if the bit at position i is set.
		if (n & (1 << i)) != 0 {
			positions = append(positions, uint64(i))
		}
	}
	return functional.Fmap(func(x uint64) Square {
		row, col := To(x)
		return Square{row, col}
	}, positions)
}

func StartBoard() Board {
	retVal := Board{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 0)
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 1)
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 2)
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 3)
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 4)
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 5)
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 6)
	retVal = SetPiece(ColoredPiece{White, Pawn}, retVal, 1, 7)
	retVal = SetPiece(ColoredPiece{White, Rook}, retVal, 0, 0)
	retVal = SetPiece(ColoredPiece{White, Knight}, retVal, 0, 1)
	retVal = SetPiece(ColoredPiece{White, Bishop}, retVal, 0, 2)
	retVal = SetPiece(ColoredPiece{White, Queen}, retVal, 0, 3)
	retVal = SetPiece(ColoredPiece{White, King}, retVal, 0, 4)
	retVal = SetPiece(ColoredPiece{White, Bishop}, retVal, 0, 5)
	retVal = SetPiece(ColoredPiece{White, Knight}, retVal, 0, 6)
	retVal = SetPiece(ColoredPiece{White, Rook}, retVal, 0, 7)

	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 0)
	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 1)
	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 2)
	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 3)
	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 4)
	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 5)
	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 6)
	retVal = SetPiece(ColoredPiece{Black, Pawn}, retVal, 6, 7)
	retVal = SetPiece(ColoredPiece{Black, Rook}, retVal, 7, 0)
	retVal = SetPiece(ColoredPiece{Black, Knight}, retVal, 7, 1)
	retVal = SetPiece(ColoredPiece{Black, Bishop}, retVal, 7, 2)
	retVal = SetPiece(ColoredPiece{Black, Queen}, retVal, 7, 3)
	retVal = SetPiece(ColoredPiece{Black, King}, retVal, 7, 4)
	retVal = SetPiece(ColoredPiece{Black, Bishop}, retVal, 7, 5)
	retVal = SetPiece(ColoredPiece{Black, Knight}, retVal, 7, 6)
	retVal = SetPiece(ColoredPiece{Black, Rook}, retVal, 7, 7)

	return retVal
}

func PawnMoves(board Board, color Color, ch chan Board) {
	homePawnRow := 1
	if color == Black {
		homePawnRow = 6
	}

	destRow := func(row int) int {
		if color == White {
			return row + 1
		} else {
			return row - 1
		}
	}
	doubleMoves := func(square Square) {
		destinationRow := 3
		midRow := 2
		if color == Black {
			midRow = 5
		}
		if color == Black {
			destinationRow = 4
		}
		if square.Row == homePawnRow && SlowPieceAt(board, Square{homePawnRow, square.Col}) == functional.Some[ColoredPiece](ColoredPiece{color, Pawn}) {
			if SlowVacantAt(board, Square{destinationRow, square.Col}) && SlowVacantAt(board, Square{midRow, square.Col}) {
				ch <- Move(board, Square{homePawnRow, square.Col}, Square{destinationRow, square.Col})
			}
		}
	}
	singleMoves := func(square Square) {
		if SlowVacantAt(board, Square{destRow(square.Row), square.Col}) {
			ch <- Move(board, Square{square.Row, square.Col}, Square{destRow(square.Row), square.Col})
		}
	}
	pawnTakeSquares := func(square Square, color Color) []Square {
		up := square.Row + 1
		if color == Black {
			up = square.Row - 1
		}
		inside := func(s Square) bool { return InsideBoard(s) }
		return functional.Filter(
			[]Square{{up, square.Col - 1}, {up, square.Col + 1}},
			inside,
		)
	}
	takes := func(square Square) {
		for _, s := range pawnTakeSquares(square, color) {
			enemy := Black
			if color == Black {
				enemy = White
			}
			if HasColorPieceAt(board, s, enemy) {
				ch <- Move(board, square, s)
			}
		}

		// use pawnTakeSquares
	}
	for _, square := range FindAll(board, Pawn, color) {
		singleMoves(square)
		doubleMoves(square)
		takes(square)
	}
	close(ch)
}

func Run() {
	defer fmt.Println("Program exiting")

	board := StartBoard()

	ch := make(chan Board)
	go PawnMoves(board, White, ch)

	for b := range ch {
		fmt.Println(b)
	}

	b := StartBoard()
	b1 := Move(b, Square{1,0}, Square{3,0})
	b2 := Move(b1, Square{6,1}, Square{4,1})
	ch2 := make(chan Board)
	go PawnMoves(b2, White, ch2)

	for b := range ch2 {
		fmt.Println(b)
	}

}
