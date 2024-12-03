package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int
	var sign string
	var num2 int

	fmt.Print("введите первое число:  ")
	fmt.Scan(&num1)
	fmt.Print("Cимволы: ")
	fmt.Scan(&sign)
	fmt.Print("введите второе число:  ")
	fmt.Scan(&num2)

	if sign == "+" {
		fmt.Println("ответ: " + strconv.Itoa(num1+num2))
	} else if sign == "-" {
		fmt.Println("ответ: " + strconv.Itoa(num1-num2))
	} else if sign == "/" {
		fmt.Println("ответ: " + strconv.Itoa(num1/num2))
	} else if sign == "*" {
		fmt.Println("ответ: " + strconv.Itoa(num1*num2))
	}
}
