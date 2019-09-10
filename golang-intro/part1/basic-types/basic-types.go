package main

import "fmt"

const pi = 3.14159 // constante. Quién lo hubiera dicho...

func main() {
	var floor int         // declaración
	floor = 13            // asignación
	var name = "John Doe" // el tipo "string" es inferido (declaración y asignación)
	active := false       // forma abreviada de declaración / asignación

	varDump(pi, floor, name, active)
}

func varDump(vars ...interface{}) {
	for i, myVar := range vars {
		fmt.Printf("The type of var %d is %T and its value %v\n", i, myVar, myVar)
	}
}
