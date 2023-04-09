package loops

import "fmt"

func For39() {
	var a, b int
	fmt.Println("Please enter A:")
	fmt.Scan(&a)
	fmt.Println("Please enter B:")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%v", i)
		}
	}

}

func For40() {
	var a, b int
	fmt.Println("Please enter A:")
	fmt.Scan(&a)
	fmt.Println("Please enter B:")
	fmt.Scan(&b)
	k := 0
	for i := a; i <= b; i++ {
		k++
		for j := 0; j < k; j++ {
			fmt.Print(i)
		}
	}

}
