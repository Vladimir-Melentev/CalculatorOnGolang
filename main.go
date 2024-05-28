package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func isRoman(n string) bool {
	_, ok := romanToArabic[n]
	return ok
}

func romanToInt(n string) int {
	return romanToArabic[n]
}

func intToRoman(n int) string {
	if n < 1 {
		panic("Result out of bounds for Roman numerals")
	}

	var result strings.Builder
	for _, entry := range arabicToRoman {
		for n >= entry.Value {
			result.WriteString(entry.Symbol)
			n -= entry.Value
		}
	}
	return result.String()
}

func parseInput(input string) (int, int, string, bool) {
	re := regexp.MustCompile(`^\s*(\d+|I{1,4}|V?I{0,3}|IX|IV|VI{0,3}|X{1,2})\s*([+\-*/])\s*(\d+|I{1,4}|V?I{0,3}|IX|IV|VI{0,3}|X{1,2})\s*$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 4 {
		panic("Invalid input format")
	}

	aStr, op, bStr := matches[1], matches[2], matches[3]

	if isRoman(aStr) && isRoman(bStr) {
		a, b := romanToInt(aStr), romanToInt(bStr)
		return a, b, op, true
	}

	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)
	if errA == nil && errB == nil {
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("Numbers out of bounds")
		}
		return a, b, op, false
	}

	panic("Invalid input format")
}

func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Invalid operation")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter calculation: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	a, b, op, isRoman := parseInput(input)
	result := calculate(a, b, op)

	if isRoman {
		if result < 1 {
			panic("Result out of bounds for Roman numerals")
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}
