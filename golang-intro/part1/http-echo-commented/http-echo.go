package main // paquete actual. Main es el punto de entrada del programa

import ( // importamos otros paquetes necesarios de la librería estándar
	"fmt"      // fmt se utiliza para formateo de datos
	"io"       // io tiene algunas funciones de entrada / salida
	"net/http" // utilidades de cliente y servidor HTTP
)

const port = 8080 // constante con el puerto en el que escuchar

func main() { // la función main del paquete main es donde comienza la ejecución del programa
	// escribimos por salida estándar (stdout) en qué puerto escucha el servidor
	fmt.Printf("Listening to requests on port %d\n", port)
	// registramos un "HTTP handler" en la ruta raíz "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "It works!") // escribimos en la Response el texto
	})

	// comenzamos a escuchar conexiones en el puerto especificado
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
