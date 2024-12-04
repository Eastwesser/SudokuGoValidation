package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

// A. Функция для удаления лишних пробелов
func removeSpaces(textLine string) string {

	for strings.Contains(textLine, "  ") {
		textLine = strings.ReplaceAll(textLine, "  ", " ")
	}

	return textLine
}

// B. Функция для смены символов вокруг "-" (и удаления "-")
func swapChars(textLine string) string {
	runes := []rune(textLine)
	var res []rune

	for i := 0; i < len(runes); i++ {
		if runes[i] == '-' && i > 0 && i < len(runes)-1 {
			res[len(res)-1], runes[i+1] = runes[i+1], res[len(res)-1]
		} else {
			res = append(res, runes[i])
		}
	}

	return string(res)
}

// C. Функция для замены "+" на "!"
func swapPlus(textLine string) string {

	for strings.Contains(textLine, "+") {
		textLine = strings.ReplaceAll(textLine, "+", "!")
	}

	return textLine
}

// D. Функция для удаления цифр и подсчета их суммы
func countSum(textLine string) string {
	var sum int
	var builder strings.Builder

	for _, char := range textLine {
		if unicode.IsDigit(char) {
			digit, _ := strconv.Atoi(string(char))
			sum += digit
		} else {
			builder.WriteRune(char)
		}
	}

	if sum > 0 {
		builder.WriteString(" ")
		builder.WriteString(strconv.Itoa(sum))

	}

	return builder.String()
}

func (t *Text) TextModifier() {

	// A
	t.Content = removeSpaces(t.Content)
	// B
	t.Content = swapChars(t.Content)
	// C
	t.Content = swapPlus(t.Content)
	// D
	t.Content = countSum(t.Content)

	fmt.Println(t.Content)
}

func main() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	if scanner.Scan() {
		text.Content = scanner.Text()
		text.TextModifier()
	} else {
		fmt.Println("Ошибка ввода:", scanner.Err())
	}
}
