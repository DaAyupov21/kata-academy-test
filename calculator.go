package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNums = [...]string{"", "I", "II", "III", "IV", "V",
	"VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV",
	"XVI", "XVII", "XVIII", "XIX", "XX"}

func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func splitTemp(temp string) (a int, b int, isRoman bool, op string, err error) {
	if strings.Contains(temp, "-") {
		return decodeTemplate(temp, "-")
	} else if strings.Contains(temp, "+") {
		return decodeTemplate(temp, "+")
	} else if strings.Contains(temp, "*") {
		return decodeTemplate(temp, "*")
	} else if strings.Contains(temp, "/") {
		return decodeTemplate(temp, "/")
	} else {
		return 0, 0, false, "", fmt.Errorf("Паника. Строка не является математической операцией.")
	}
}

func decodeTemplate(temp string, op string) (a int, b int, isRoman bool, operate string, err error) {
	val := strings.Split(temp, " "+op+" ")
	if len(val) > 2 {
		return 0, 0, false, "", fmt.Errorf("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	aIsRoman, a := isRomanNum(val[0])
	bIsRoman, b := isRomanNum(val[1])
	operate = op
	if aIsRoman != bIsRoman {
		return 0, 0, false, "", fmt.Errorf("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	if aIsRoman && a <= b {
		return 0, 0, false, "", fmt.Errorf("Выдача паники, в римской системе счисления нет отрицательных чисел и 0.")
	}
	if a < 1 || b < 1 && a > 10 || b > 10 {
		return 0, 0, false, "", fmt.Errorf("Выдача паники, операции проводятся только с натуральными числами, но не больше 10")
	}
	return a, b, aIsRoman, operate, nil
}

func operate(a int, b int, op string) (int, error) {
	if op == "+" {
		return a + b, nil
	} else if op == "-" {
		return a - b, nil
	} else if op == "*" {
		return a * b, nil
	} else if op == "/" {
		return a / b, nil
	} else {
		return -1, fmt.Errorf("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
}

func convertIntToRoman(number int) string {
	var tens = [...]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	var units = romanNums[:9]
	return tens[(number%100)/10] + units[number%10]
}

func isRomanNum(a string) (bool, int) {

	for i, v := range romanNums {
		if a == v {
			return true, i
		}
	}
	aNum, _ := strconv.ParseInt(a, 10, 64)
	return false, int(aNum)
}

func main() {
	var temp string
	temp = scan()
	a, b, isRoman, op, err := splitTemp(temp)
	if err != nil {
		println(err.Error())
		return
	}
	result, err := operate(a, b, op)
	if err != nil {
		println(err.Error())
	} else {
		if isRoman {
			println(convertIntToRoman(result))
		} else {
			println(result)
		}

	}

}
