package main

import "fmt"

/*
board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

boardTrue := [9][9]int{
{"5","3",".",".","7",".",".",".","."},
{"6",".",".","1","9","5",".",".","."},
{".","9","8",".",".",".",".","6","."},
{"8",".",".",".","6",".",".",".","3"},
{"4",".",".","8",".","3",".",".","1"},
{"7",".",".",".","2",".",".",".","6"},
{".","6",".",".",".",".","2","8","."},
{".",".",".","4","1","9",".",".","5"},
{".",".",".",".","8",".",".","7","9"},
}


boardFalse := [9][9]int{
{"8","3",".",".","7",".",".",".","."},
{"6",".",".","1","9","5",".",".","."},
{".","9","8",".",".",".",".","6","."},
{"8",".",".",".","6",".",".",".","3"},
{"4",".",".","8",".","3",".",".","1"},
{"7",".",".",".","2",".",".",".","6"},
{".","6",".",".",".",".","2","8","."},
{".",".",".","4","1","9",".",".","5"},
{".",".",".",".","8",".",".","7","9"},
}

func isValidSudoku(board [][]byte) bool {

}