package chess

import (
	"fmt"
	"go-poc/functional"
)

type Piece int

const (
	WhitePawn Piece = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

func (p Piece) String() string {
	switch p {
	case WhitePawn:
		return "♙"
	case BlackPawn:
		return "♟️"
	case WhiteKnight:
		return "♘"
	case BlackKnight:
		return "♞"
	case WhiteBishop:
		return "♗"
	case BlackBishop:
		return "♝"
	case WhiteRook:
		return "♖"
	case BlackRook:
		return "♖"
	case WhiteQueen:
		return "♕"
	case BlackQueen:
		return "♛"
	case WhiteKing:
		return "♔"
	case BlackKing:
		return "♚"
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
			if (functional.Empty[Piece](pieceMaybe)) {
				if (col + row) % 2 == 0 {
					retVal += "▫"
				} else {
					retVal += "▪"
				}
			} else {
				retVal += functional.Get[Piece](pieceMaybe).String()
			}
		}
		retVal += "\n"
	}
	return retVal
}

// set a piece at the board
func SetPiece(piece Piece, board Board, row, col int) Board {
	board[int(piece)] = setBit(board[int(piece)], row, col)
	return board
}

// remove a piece from the board
func ClearPiece(board Board, row, col int) {
	for pi := 0; pi < 12; pi++ {
		clearBit(board[pi], row, col)
	}
}

// setBit sets the bit at the given position on the chessboard.
func setBit(board uint64, row, col int) uint64 {
	board |= 1 << (row*8 + col)
	return board
}

// clearBit clears the bit at the given position on the chessboard.
func clearBit(board uint64, row, col int) {
	board &^= 1 << (row*8 + col)
}

// isSet checks if the bit at the given position on the chessboard is set.
func isSet(board uint64, row, col int) bool {
	return board&(1<<(row*8+col)) != 0
}

func PieceAt(board Board, row, col int) functional.Maybe[Piece] {
	for pi := 0; pi < 12; pi++ {
		if isSet(board[pi], row, col) {
			return functional.Some(Piece(pi))
		}
	}
	return functional.None[Piece]()
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
	retVal = SetPiece(WhitePawn, retVal, 1, 0)
	retVal = SetPiece(WhitePawn, retVal, 1, 1)
	retVal = SetPiece(WhitePawn, retVal, 1, 2)
	retVal = SetPiece(WhitePawn, retVal, 1, 3)
	retVal = SetPiece(WhitePawn, retVal, 1, 4)
	retVal = SetPiece(WhitePawn, retVal, 1, 5)
	retVal = SetPiece(WhitePawn, retVal, 1, 6)
	retVal = SetPiece(WhitePawn, retVal, 1, 7)
	retVal = SetPiece(WhiteRook, retVal, 0, 0)
	retVal = SetPiece(WhiteKnight, retVal, 0, 1)
	retVal = SetPiece(WhiteBishop, retVal, 0, 2)
	retVal = SetPiece(WhiteQueen, retVal, 0, 3)
	retVal = SetPiece(WhiteKing, retVal, 0, 4)
	retVal = SetPiece(WhiteBishop, retVal, 0, 5)
	retVal = SetPiece(WhiteKnight, retVal, 0, 6)
	retVal = SetPiece(WhiteRook, retVal, 0, 7)
	retVal = SetPiece(BlackPawn, retVal, 6, 0)
	retVal = SetPiece(BlackPawn, retVal, 6, 1)
	retVal = SetPiece(BlackPawn, retVal, 6, 2)
	retVal = SetPiece(BlackPawn, retVal, 6, 3)
	retVal = SetPiece(BlackPawn, retVal, 6, 4)
	retVal = SetPiece(BlackPawn, retVal, 6, 5)
	retVal = SetPiece(BlackPawn, retVal, 6, 6)
	retVal = SetPiece(BlackPawn, retVal, 6, 7)
	retVal = SetPiece(BlackRook, retVal, 7, 0)
	retVal = SetPiece(BlackKnight, retVal, 7, 1)
	retVal = SetPiece(BlackBishop, retVal, 7, 2)
	retVal = SetPiece(BlackQueen, retVal, 7, 3)
	retVal = SetPiece(BlackKing, retVal, 7, 4)
	retVal = SetPiece(BlackBishop, retVal, 7, 5)
	retVal = SetPiece(BlackKnight, retVal, 7, 6)
	retVal = SetPiece(BlackRook, retVal, 7, 7)

	return retVal
}

func Run() {
	fmt.Println(WhitePawn)
	start := StartBoard()
	fmt.Println(start)
}
