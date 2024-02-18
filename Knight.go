package main

import (
	"fmt"
	"unicode"
)

const N = 8

var (
	dx = [8]int{2, 1, -1, -2, -2, -1, 1, 2}
	dy = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
)

type Position struct {
	x, y int
}

func toBoardCoordinates(pos string) (int, int, error) {
	if len(pos) != 2 {
		return 0, 0, fmt.Errorf("invalid position")
	}

	column := unicode.ToUpper(rune(pos[0]))
	row := pos[1] - '1'

	if column < 'A' || column > 'H' || row < 0 || row >= N {
		return 0, 0, fmt.Errorf("position out of bounds")
	}

	return int(column - 'A'), int(row), nil
}

func toChessNotation(x, y int) string {
	return fmt.Sprintf("%c%d", 'A'+x, y+1)
}

func isSafe(x, y int, board [N][N]int) bool {
	return x >= 0 && y >= 0 && x < N && y < N && board[x][y] == -1
}

func solveKTUtil(board *[N][N]int, currPos Position, movei int, moves *[]string) bool {
	if movei == N*N {
		return true
	}

	for k := 0; k < 8; k++ {
		nextX := currPos.x + dx[k]
		nextY := currPos.y + dy[k]
		if isSafe(nextX, nextY, *board) {
			(*board)[nextX][nextY] = movei
			*moves = append(*moves, fmt.Sprintf("%s ---> %s", toChessNotation(currPos.x, currPos.y), toChessNotation(nextX, nextY)))
			if solveKTUtil(board, Position{nextX, nextY}, movei+1, moves) {
				return true
			} else {
				(*board)[nextX][nextY] = -1 // backtracking
				*moves = (*moves)[:len(*moves)-1]
			}
		}
	}
	return false
}

func solveKT(startPos string) bool {
	startX, startY, err := toBoardCoordinates(startPos)
	if err != nil {
		fmt.Println(err)
		return false
	}

	board := [N][N]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[i][j] = -1
		}
	}

	board[startX][startY] = 0
	var moves []string

	if !solveKTUtil(&board, Position{startX, startY}, 1, &moves) {
		fmt.Println("Solution does not exist")
		return false
	} else {
		for _, move := range moves {
			fmt.Println(move)
		}
		fmt.Printf("Sequence finished in %d steps\n", len(moves))
	}
	return true
}

func main() {
	var startPos string
	fmt.Println("Enter the starting position (e.g., A1, H5, B6):")
	fmt.Scan(&startPos)

	if !solveKT(startPos) {
		fmt.Println("No solution found")
	}
}
