package INTEGER

import "fmt"

func Integer16() {

	var n int
	fmt.Println("Enter some number XXX: ")
	fmt.Scan(&n)
	s := n % 100
	m := n - s
	l := s % 10
	b := s / 10
	a := m + (l * 10) + b
	fmt.Printf("Итог: %v", a)
}

func Integer30() {
	var a int
	fmt.Println("Год:")
	fmt.Scan(&a)
	if a%100 != 0 {
		a /= 100
		a++
		fmt.Printf("Век: %v", a)
	} else {
		a /= 100
		fmt.Printf("Век: %v", a)
	}

}

func Integer23() {

	var s int
	fmt.Println("Enter some amount of sec:")
	fmt.Scan(&s)
	l := (s / 60) % 60
	fmt.Printf("Amount of minutes of the last hour: %v", l)
}
