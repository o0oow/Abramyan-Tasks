package main

import (
	"awesomeProject/ARRAY"
	"awesomeProject/BOOLEAN"
	"awesomeProject/IF"
	"awesomeProject/INTEGER"
	"awesomeProject/begin"
	"awesomeProject/loops"
	"fmt"
)

func main() {

	var w int
	fmt.Print("Tasks from Abramyan book:")
	fmt.Print("\n1.Begin31 2.Begin36 3.Begin38 \n4.Integer16 5.Integer30 6.Integer23\n ")
	fmt.Print("7.Boolean18 8.Boolean22 9.Boolean35\n 10.If28 11.If29 12.If30\n 13.For39 14.For40\n")
	fmt.Print("15.Array2 16.Array14 17.Array21\n")

	fmt.Println("Enter the number of the task:")
	fmt.Scanln(&w)

	switch w {
	case 1:
		begin.Begin31()
	case 2:
		begin.Begin36()
	case 3:
		begin.Begin38()
	case 4:
		INTEGER.Integer16()
	case 5:
		INTEGER.Integer30()
	case 6:
		INTEGER.Integer23()
	case 7:
		BOOLEAN.Boolean18()
	case 8:
		BOOLEAN.Boolean22()
	case 9:
		BOOLEAN.Boolean35()
	case 10:
		IF.If28()
	case 11:
		IF.If29()
	case 12:
		IF.If30()
	case 13:
		loops.For39()
	case 14:
		loops.For40()
	case 15:
		ARRAY.Array2()
	case 16:
		ARRAY.Array14()
	case 17:
		ARRAY.Array21()
	default:
		fmt.Println("Here is no such task ")

	}
}
