package main

import "fmt"

func main() {
	// START OMIT
	nums := []int{1, 2, 3}

	func() {
		for _, num := range nums { // nums está definido fuera de la función
			fmt.Println(num)
		}
	}()
	// END OMIT
}
