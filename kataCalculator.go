package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var addition = "+"
var subtraction = "-"
var multiplication = "*"
var division = "/"
var result int
var resultA bool
var resultB bool
var resultC bool
var resultD bool
var i string
var number1 string
var number2 string
var number3 int
var number4 int

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите пример...")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	textLength := len(text)

	resultA = strings.Contains(text, addition)
	resultB = strings.Contains(text, subtraction)
	resultC = strings.Contains(text, multiplication)
	resultD = strings.Contains(text, division)

	switch {
	case resultA == true:
		i = "+"
	case resultB == true:
		i = "-"
	case resultC == true:
		i = "*"
	case resultD == true:
		i = "/"
	}

	if textLength < 3 {
		err := errors.New("ошибка: введите пример с двумя операндами и одним оператором")
		fmt.Println(err)
		os.Exit(0)
	}

	if textLength > 5 {
		err := errors.New("ошибка: введите пример с двумя операндами и одним оператором")
		fmt.Println(err)
		os.Exit(0)
	}

	numbers := strings.Split(text, i)
	number1 = numbers[0]
	number2 = numbers[1]

	roman := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	rNumber1, ok := roman[number1]
	rNumber2, status := roman[number2]

	if ok && status {
		number3 = rNumber1
		number4 = rNumber2
	} else {
		number3, _ = strconv.Atoi(number1)
		number4, _ = strconv.Atoi(number2)
	}

	if ok != status {
		err := errors.New("ошибка: введены операнды разных систем счисления")
		fmt.Println(err)
		os.Exit(0)
	}

	if len(numbers) < 2 {
		err := errors.New("ошибка: введите два операнда и оператор между ними")
		fmt.Println(err)
		os.Exit(0)
	}

	if len(numbers) > 2 {
		err := errors.New("ошибка: введите два операнда и оператор между ними")
		fmt.Println(err)
		os.Exit(0)
	}

	remainder := strings.Contains(text, ",")
	remainder2 := strings.Contains(text, ".")

	if remainder == true {
		err := errors.New("ошибка: введите целое число")
		fmt.Println(err)
		os.Exit(0)
	}

	if remainder2 == true {
		err := errors.New("ошибка: введите целое число")
		fmt.Println(err)
		os.Exit(0)
	}

	if number3 > 10 {
		err := errors.New("ошибка: введите числа в диапазоне от 1 до 10")
		fmt.Println(err)
		os.Exit(0)
	}

	if number4 > 10 {
		err := errors.New("ошибка: введите числа в диапазоне от 1 до 10")
		fmt.Println(err)
		os.Exit(0)
	}

	if number3 < 1 {
		err := errors.New("ошибка: введите числа в диапазоне от 1 до 10")
		fmt.Println(err)
		os.Exit(0)
	}

	if number4 < 1 {
		err := errors.New("ошибка: введите числа в диапазоне от 1 до 10")
		fmt.Println(err)
		os.Exit(0)
	}

	if resultA == false {
		if resultB == false {
			if resultC == false {
				if resultD == false {
					err := errors.New("ошибка: введите пример для решения")
					fmt.Println(err)
					os.Exit(0)
				}
			}
		}
	}

	switch {
	case i == "+":
		result = number3 + number4
	case i == "-":
		result = number3 - number4
	case i == "*":
		result = number3 * number4
	case i == "/":
		result = number3 / number4
	}

	arabic := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX",
		20: "XX", 21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX",
		30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX",
		40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX",
		50: "L", 51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX",
		60: "LX", 61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX",
		70: "LXX", 71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX",
		80: "LXXX", 81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX",
		90: "XC", 91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVIII", 98: "XCVIII", 99: "XCIX",
		100: "C",
	}

	resultNew := strconv.Itoa(result)
	aNumber, _ := arabic[result]

	if ok && status {
		if result <= 0 {
			err := errors.New("ошибка: результат меньше либо равен нулю")
			fmt.Println(err)
			os.Exit(0)
		}
	}

	if ok && status {
		resultNew = aNumber
		fmt.Println(resultNew)
	} else {
		fmt.Println(result)
	}
}
