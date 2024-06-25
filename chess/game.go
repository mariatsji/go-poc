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
			return "♖"
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
			pieceMaybe := PieceAt(b, row, col)
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

// set a piece at the board
func SetPiece(piece ColoredPiece, board Board, row, col int) Board {
	board[int(piece.Piece)] = setBit(board[int(piece.Piece)], row, col)
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
func Move(board Board, fromRow, fromCol, toRow, toCol int) Board {
	pieceMaybe := PieceAt(board, fromRow, fromCol)
	functional.DoIf(pieceMaybe,
		(func(coloredPiece ColoredPiece) {
			board = SetPiece(coloredPiece, board, toRow, toCol)
		}))
	board = ClearPiece(board, fromRow, fromCol)
	return board
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

func PieceAt(board Board, row, col int) functional.Maybe[ColoredPiece] {
	for pi := 0; pi < 12; pi++ {
		color := White
		if pi > 5 {
			color = Black
		}
		if isSet(board[pi], row, col) {
			return functional.Some(ColoredPiece{color, Piece(pi)})
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

func PawnMoves(color Color, board Board) []Board {
	twoSteps := make([]Board, 0)
	homePawnRow := 1
	if color == Black {
		homePawnRow = 6
	}
	destinationRow := 3
	if color == Black {
		destinationRow = 4
	}
	for col := 0; col < 7; col++ {
		if PieceAt(board, homePawnRow, col) == functional.Some[ColoredPiece](ColoredPiece{color, Pawn}) {
			twoSteps = append(twoSteps, Move(board, col, homePawnRow, col, destinationRow))
		}
	}
	oneSteps := make([]Board, 0)
	takesRight := make([]Board, 0)
	takesLeft := make([]Board, 0)
	promotions := make([]Board, 0)
	promotionsTakesLeft := make([]Board, 0)
	promotionsTakesRight := make([]Board, 0)
	enPassant := make([]Board, 0)
	retVal :=
		append(
			append(
				append(
					append(
						append(
							append(
								append(
									twoSteps,
									oneSteps...),
								takesRight...),
							takesLeft...),
						promotions...),
					promotionsTakesLeft...),
				promotionsTakesRight...),
			enPassant...)

	return retVal
}

func Run() {
	board := StartBoard()
	fmt.Println(board)

	board = Move(board, 1, 4, 3, 4)
	fmt.Println(board)
}
