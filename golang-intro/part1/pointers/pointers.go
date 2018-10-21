package main

import "fmt"

func main() {
	//START OMIT
	type user struct {
		Name string
	}

	u1 := user{"Marcos"}
	u2 := u1
	u3 := &u1 // u3 apunta al valor u1 (Marcos)
	u2.Name = "Mari Carmen"
	u3.Name = "Jacinto" // cambiamos el valor original

	fmt.Println(u1, u2, u3)
	// END OMIT
}
