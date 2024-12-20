package main

import (
	"fmt"
)

type Text struct {
	Content string
}

// =====================================================================================================================
func checkRowIsValid(board []byte) bool {
	var res bool
	// Просто итерируемся по всем девятерым массивам и ищем там уникальные строки. Точки игнорим.
	// допустим:
	//{"5", "3", ".", ".", "7", ".", ".", ".", "."}, - это первый массив, ищем внутри него уникальные значения
	//{"6", ".", ".", "1", "9", "5", ".", ".", "."}, - это второй массив, ищем внутри него уникальные значения
	//{".", "9", "8", ".", ".", ".", ".", "6", "."}, - это третий массив, ищем внутри него уникальные значения
	//{"8", ".", ".", ".", "6", ".", ".", ".", "3"}, - это четвертый массив, ищем внутри него уникальные значения
	//{"4", ".", ".", "8", ".", "3", ".", ".", "1"}, - это пятый массив, ищем внутри него уникальные значения
	//{"7", ".", ".", ".", "2", ".", ".", ".", "6"}, - это шестой массив, ищем внутри него уникальные значения
	//{".", "6", ".", ".", ".", ".", "2", "8", "."}, - это седьмой массив, ищем внутри него уникальные значения
	//{".", ".", ".", "4", "1", "9", ".", ".", "5"}, - это восьмой массив, ищем внутри него уникальные значения
	//{".", ".", ".", ".", "8", ".", ".", "7", "9"}, - это девятый массив, ищем внутри него уникальные значения

	for i := 0; i < len(board); i++ {
		for j := i + 1; j < len(board); j++ {
			if board[i] == board[j] {
				return false
			}
		}
	}

	return res
}

// =====================================================================================================================
func checkColumnIsValid(board [][]byte) bool {
	// Итерируемся по [0][0]; [1][0]; [2][0]; [3][0]; [4][0]; [5][0]; [6][0]; [7][0]; [8][0] в каждом массиве берем первые значения
	// Итерируемся по [0][1]; [1][1]; [2][1]; [3][1]; [4][1]; [5][1]; [6][1]; [7][1]; [8][1] в каждом массиве берем вторые значения
	// Итерируемся по [0][2]; [1][2]; [2][2]; [3][2]; [4][2]; [5][2]; [6][2]; [7][2]; [8][2] в каждом массиве берем третьи значения
	// Итерируемся по [0][3]; [1][0]; [2][0]; [3][0]; [4][0]; [5][0]; [6][0]; [7][0]; [8][0] в каждом массиве берем четвертые значения
	// Итерируемся по [0][4]; [1][1]; [2][1]; [3][1]; [4][1]; [5][1]; [6][1]; [7][1]; [8][1] в каждом массиве берем пятые значения
	// Итерируемся по [0][5]; [1][2]; [2][2]; [3][2]; [4][2]; [5][2]; [6][2]; [7][2]; [8][2] в каждом массиве берем шестые значения
	// Итерируемся по [0][6]; [1][0]; [2][0]; [3][0]; [4][0]; [5][0]; [6][0]; [7][0]; [8][0] в каждом массиве берем седьмые значения
	// Итерируемся по [0][7]; [1][1]; [2][1]; [3][1]; [4][1]; [5][1]; [6][1]; [7][1]; [8][1] в каждом массиве берем восьмые значения
	// Итерируемся по [0][8]; [1][2]; [2][2]; [3][2]; [4][2]; [5][2]; [6][2]; [7][2]; [8][2] в каждом массиве берем девятые значения
	// For Go specifically, you can use a map[rune]bool or map[string]bool to track seen digits.

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if checkRowIsValid(board[i]) {
				return false
			}
		}
	}
	return true
}

// =====================================================================================================================
func checkSubGridIsValid(board [][]byte) bool {
	return true
}

// =====================================================================================================================
func (board *Text) isValidSudoku() {

	// ARE ROWS VALID?
	board.Content = checkRowIsValid(board.Content)
	// ARE COLUMNS VALID?
	board.Content = checkColumnIsValid(board.Content)
	// ARE GRIDS VALID?
	board.Content = checkSubGridIsValid(board.Content)

	fmt.Println(board.Content)

}

// =====================================================================================================================
func main() {
	// передаем уже приготовленные судоку на проверку
	boardTrue := [9][9]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	boardFalse := [9][9]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	fmt.Println("Это Судоку верное.")
	fmt.Println(isValidSudoku(boardTrue))
	fmt.Println("Это Судоку неверное!!!")
	fmt.Println(isValidSudoku(boardFalse))

	// Далее вариант кода, когда мы хотим ввести новый вид судоку, через консоль. Думаю bufio отлично подойдет для ввода.
	//board := &Text{}
	//scanner := bufio.NewScanner(os.Stdin)
	//
	//fmt.Println("Please enter a valid Sudoku input: ")
	//if scanner.Scan() {
	//	board.Content = scanner.Text()
	//	board.isValidSudoku()
	//} else {
	//	fmt.Println("Please enter a valid Sudoku input: ", scanner.Err())
	//}
}
