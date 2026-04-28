package main

import (
	"fmt"
	"math/rand" // libreria nativa de go para generar numeros aleatorios
)

func AdivinaNumero() {
	// Generamos un numero aleatorio pero que se entre 1 y 100
	// rand.Intn(100) me da un número del 0 al 99, por eso le sumo 1 al final.
	secreto := rand.Intn(100) + 1

	intentos := 5
	var entrada int

	fmt.Println("¡Bienvenido! He pensado en un número oculto entre 1 y 100.")

	// Creo un ciclo que se repetirá únicamente mientras queden vidas.
	for intentos > 0 {
		
		// Imprimo cuántos intentos quedan y pido el número.
		fmt.Printf("Tienes %d intentos. Ingresa tu número: ", intentos)
		
		// Pauso el programa para leer lo que escribe el usuario. 
		// El '&' guarda el número directamente en la memoria de mi variable 'entrada'.
		fmt.Scan(&entrada)

		// Evaluamos la condición de victoria primero.
		if entrada == secreto {
			fmt.Println("¡Correcto! Has adivinado el número.")
			return
		}

		// Si no ganó, evalúo si le doy la pista de "menor" o "mayor".
		if entrada > secreto {
			fmt.Println("Pista: El número secreto es menor.")
		} else {
			fmt.Println("Pista: El número secreto es mayor.")
		}

		// Como falló, le resto un intento a mi contador.
		intentos--
	}

	// Si el ciclo terminó por completo y nunca se activó el 'return' de victoria, 
	// significa que se acabaron los intentos y el usuario perdió.
	fmt.Printf("Has perdido. El número era %d.\n", secreto)
}

func main() {
	AdivinaNumero()

}