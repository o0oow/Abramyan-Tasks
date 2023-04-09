package ARRAY

import "fmt"

func Array21() {
	var n, v int
	fmt.Println("Enter the size of array:")
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		arr = append(arr, v)

	}

	for i := 0; i < n; i++ {
		fmt.Printf("%v", arr[i])
	}

	var l, k int
	sum := 0
	fmt.Println("\nEnter K:")
	fmt.Scan(&k)
	fmt.Println("\nEnter L:")
	fmt.Scan(&l)
	for i := k; i <= l; i++ {
		sum += arr[i]
	}
	mean := sum / (l - k + 1)
	fmt.Printf("Here is mean: %v", float64(mean))
}

func Array2() {
	var arr []int
	var n int
	fmt.Println("Enter N:")
	fmt.Scan(&n)
	for i := 2; len(arr) < n; i *= 2 {
		arr = append(arr, i)
	}
	fmt.Println(arr)
}

func Array14() {
	var n, v int
	fmt.Println("Enter the size of array:")
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		arr = append(arr, v)
	}

	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], "")
	}
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], "")
	}

}
