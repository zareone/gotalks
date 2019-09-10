package main

import (
	"fmt"
)

func main() {
	// START OMIT
	// inicialización y asignación de valores (notación literal)
	machines := map[string]int{
		"bassline": 303,
	}
	machines["drumbox"] = 808 // inserción de una clave

	fmt.Printf("machines: %v\n", machines)
	delete(machines, "drumbox")              // borrado de una clave
	model := machines["drumbox"]             // acceso a la clave borrada
	fmt.Printf("drumbox model: %v\n", model) // valor "zero" al no existir

	model, ok := machines["drumbox"] // segundo parámetro para comprobar existencia de la clave
	fmt.Printf("drumbox model: %v, ok: %v\n", model, ok)
	fmt.Printf("machines after delete: %v\n", machines)
	// END OMIT
}
