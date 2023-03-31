package BOOLEAN

import "fmt"

func Boolean18() {
	var a, b, c int
	var B bool
	fmt.Println("Enter integer numbers:")
	fmt.Scan(&a, &b, &c)
	if a == b || b == c || c == a {
		B = true
	}
	fmt.Printf("Среди трех данных целых чисел есть хотя бы одна пара совпадающих:%v", B)
}

func Boolean22() {
	var n int
	var B bool
	fmt.Println("Enter XXX number:")
	fmt.Scan(&n)
	a := n / 100
	b := (n % 100) / 10
	c := n % 10
	if a > b && b > c {
		B = true
	} else if a < b && b < c {
		B = true
	}

	fmt.Printf("Цифры данного числа образуют возрастающую или убывающую последовательность: %v", B)
}

func Boolean35() {
	var x1, y1, x2, y2 int
	var B bool
	fmt.Println(" 1 координаты шахматной доски(1-8):")
	fmt.Scan(&x1, &y1)
	fmt.Println(" 2 координаты шахматной доски(1-8):")
	fmt.Scan(&x2, &y2)
	if (x1+y1)%2 == 0 && (x2+y2)%2 == 0 {
		B = true
	}
	if ((x1+y1)%2 == 1) && ((x2+y2)%2 == 1) {
		B = true
	}
	fmt.Printf("«Данные поля имеют одинаковый цвет:%v", B)

}
