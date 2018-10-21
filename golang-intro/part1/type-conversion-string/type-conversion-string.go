package main

import (
	"fmt"
	"strconv"
)

func main() {
	// START OMIT
	i := 8
	s := string(i) // nope!
	s2 := strconv.Itoa(i)

	fmt.Printf("i: %v, s: %v, s2: %v", i, s, s2)
	// END OMIT

}
