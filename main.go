package main

import (
	"bufio"
	"fmt"
	"os"
)

type Sudoku struct {
	Content bool
}

func checkRowIsValid() bool {
	return true
}

func checkColumnIsValid() bool {
	return true
}

func checkSubGridIsValid() bool {
	return true
}

func (board *Sudoku) isValidSudoku() {

	// ARE ROWS VALID?
	board.Content = checkRowIsValid(board.Content)
	// ARE COLUMNS VALID?
	board.Content = checkColumnIsValid(board.Content)
	// ARE GRIDS VALID?
	board.Content = checkSubGridIsValid(board.Content)

	fmt.Println(board.Content)

}

func main() {
	board := &Sudoku{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter a valid Sudoku input: ")
	if scanner.Scan() {
		board.Content = scanner.Text()
		board.isValidSudoku()
	} else {
		fmt.Println("Please enter a valid Sudoku input: ", scanner.Err())
	}
}

//boardTrue := [9][9]int{
//{"5", "3", ".", ".", "7",".", ".", ".", "."},
//{"6", ".", ".", "1", "9","5", ".", ".", "."},
//{".", "9", "8", ".", ".",".", ".", "6", "."},
//{"8", ".", ".", ".", "6",".", ".", ".", "3"},
//{"4", ".", ".", "8", ".","3", ".", ".", "1"},
//{"7", ".", ".", ".", "2",".", ".", ".", "6"},
//{".", "6", ".", ".", ".",".", "2", "8", "."},
//{".", ".", ".", "4", "1","9", ".", ".", "5"},
//{".", ".", ".", ".", "8",".", ".", "7", "9"},
//}
//
//
//boardFalse := [9][9]int{
//{"8", "3", ".", ".", "7",".", ".", ".", "."},
//{"6", ".", ".", "1", "9","5", ".", ".", "."},
//{".", "9", "8", ".", ".",".", ".", "6", "."},
//{"8", ".", ".", ".", "6",".", ".", ".", "3"},
//{"4", ".", ".", "8", ".","3", ".", ".", "1"},
//{"7", ".", ".", ".", "2",".", ".", ".", "6"},
//{".", "6", ".", ".", ".",".", "2", "8", "."},
//{".", ".", ".", "4", "1","9", ".", ".", "5"},
//{".", ".", ".", ".", "8",".", ".", "7", "9"},
//}
