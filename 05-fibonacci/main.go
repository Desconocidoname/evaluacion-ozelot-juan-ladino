package main

import (
	"fmt" // imprimir y leer en la consola
	"strings" // limpiar el formato final al imprimir
)

/// Al poner '(s []int)' en la firma, estoy usando un "retorno nombrado"
// GHo nos ayuda  a crear automaticamente el arreglo vacio 's'

func GenerarFibonacci(n int) (s []int) {

	// Creo dos variables 'a' (empieza en 0) y 'b' (empieza en 1).
	// Condición: El ciclo se repite mientras el tamaño de mi arreglo sea menor a 'n'.
	// Avance paralelo: En cada vuelta, 'a' toma el valor de 'b', y al mismo tiempo 'b' toma la suma de 'a+b'.
	for a, b := 0, 1; len(s) < n; a, b = b, a+b {
		
		// En cada vuelta, solo agarro el valor actual de 'a' y lo meto a mi arreglo.
		s = append(s, a)
	}

	// Como usé un retorno nombrado, solo tengo que devolver 's' ya lleno.
	return s
}

func main() {
	var posiciones int

	fmt.Print("Ingresa la posición (n) para la sucesión de Fibonacci: ")

	fmt.Scan(&posiciones)

	// si me piden 0 o números negativos, detengo el programa
	if posiciones <= 0 {
		fmt.Println("Por favor ingresa un número mayor a 0.")
		return
	}

	resultado := GenerarFibonacci(posiciones)

	textoLimpio := strings.Trim(fmt.Sprint(resultado), "[]")

	fmt.Printf("\nSi n = %d, se debe mostrar la siguiente serie\n", posiciones)
	fmt.Println("\n", textoLimpio)
}