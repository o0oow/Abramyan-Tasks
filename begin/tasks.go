package begin

import "fmt"

func Begin36() {
	var v1, v2, T float32
	fmt.Println("Скорость 1-ой машины(км/ч):")
	fmt.Scan(&v1)
	fmt.Println("Скорость 2-ой машины(км/ч):")
	fmt.Scan(&v2)
	fmt.Println("Время(ч):")
	fmt.Scan(&T)
	S := (v1 + v2) * T
	fmt.Printf("Общий путь(км): %v", S)
}

func Begin38() {
	var A, B float64
	fmt.Println("Ввелите A:")
	fmt.Scan(&A)
	fmt.Println("Ввелите B:")
	fmt.Scan(&B)
	x := (-1 * B) / A
	fmt.Printf("X равен: %v", x)
}

func Begin31() {
	var Tf float32
	fmt.Printf("Введите количество градусов по Фарейнгейт:")
	fmt.Scan(&Tf)
	Tc := (Tf - 32) * 5 / 9
	fmt.Printf(" Цельсия: %v", Tc)
}
