package IF

import "fmt"

func If28() {
	var a int
	fmt.Println("Enter the year u want(XXXX):")
	fmt.Scan(&a)
	if a%400 == 0 || a%4 == 0 && a%100 != 0 {
		fmt.Println("In this year there are 366 days")
		If28()
	} else {
		fmt.Println("In this year there are 365 days")
		If28()
	}
}

func If29() {
	var a int
	fmt.Println("Type some number:")
	fmt.Scan(&a)

	if a < 0 && a%2 == 0 {
		fmt.Println("This is a negative and even number!!")
	} else if a < 0 && a%2 != 0 {
		fmt.Println("This is a negative and odd number!!")
	} else if a == 0 {
		fmt.Println("This is a null number!!")
	} else if a > 0 && a%2 != 0 {
		fmt.Println("This is a positive and odd number!!")
	} else if a > 0 && a%2 == 0 {
		fmt.Println("This is a positive and even number!!")
	}
}

func If30() {
	var a int
	fmt.Println("Type some number (1-999):")
	fmt.Scan(&a)
	if a >= 100 && a%2 == 0 {
		fmt.Println("This is three-digit and even number!")
		If30()
	} else if a >= 100 && a%2 != 0 {
		fmt.Println("This is three-digit and odd number!")
		If30()
	} else if a < 100 && a%2 == 0 {
		fmt.Println("This is two-digit and even number!")
		If30()
	} else if a < 100 && a%2 != 0 {
		fmt.Println("This is two-digit and odd number!")
		If30()
	}

}
