package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Input harus bilangan positif")
	} else {
		currentNum := 1

		for i := 1; i <= n; i++ {
			fmt.Print(currentNum)
			if i != n {
				fmt.Print("-")
			}
			currentNum += i
		}

		fmt.Println()
	}
}
