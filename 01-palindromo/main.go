package main

import (
	"bufio"   // nos permite leer textos largos o con espacios.
	"fmt"     // imprimir en consola.
	"os"      // Librería para acceder al sistema operativo.
	"strings"
	"unicode"
)

// ValidarPalindromo recibimos un String devuelve un String con el resultado
func ValidarPalindromo(s string) string {
	
	// Creo 'l', que es un 'slice' (como un array en JS).
	// Con 'rune' llamamos a un solo caracter.
	var l []rune
	
	// El 'for' recorre el texto. 'strings.ToLower(s)' convierte todo a minúsculas primero.
	// 'range' devuelve el índice y la letra. Usamos '_' para ignorar el índice que no nos sirve.
	// 'r' guarda la letra que estamos analizando en esa vuelta del ciclo.
	for _, r := range strings.ToLower(s) {
		
		// Verificamos si la letra actual 'r' es realmente una letra o un número (ignoramos espacios y signos).
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			// Si es válido, lo agregamos ('append') a nuestro array 'l' (como hacer l.push(r) en JS).
			l = append(l, r)
		}
	}
	
	// Revisamos la mitad de la palabra (len(l)/2).
	for i := 0; i < len(l)/2; i++ {
		
		// Comparamos la letra de la izquierda (l[i]) con la de la derecha (l[len(l)-1-i]).
		// Si en algún momento NO son iguales (!=), cortamos la función aquí mismo.
		if l[i] != l[len(l)-1-i] {
			return "Inválido"
		}
	}
	
	// Si el ciclo terminó y nunca encontró diferencias, significa que sí es un palíndromo.
	return "Válido"
}

func main() {
	
	fmt.Print("Ingresa una palabra o frase para validar: ")
	// (Standard Input), que es la entrada por defecto de la terminal.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	// scanner.Text() extrae exactamente la frase completa que el usuario escribió.
	entradaUsuario := scanner.Text()

	resultado := ValidarPalindromo(entradaUsuario)

	fmt.Printf("Resultado: La frase '%s' es %s\n", entradaUsuario, resultado)
}