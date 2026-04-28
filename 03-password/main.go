package main

import (
	"bufio"   
	"fmt"     
	"os"      // Para acceder a la entrada de la terminal
	"strings" // Para buscar caracteres o espacios en el texto
	"unicode" // Para saber si una letra es mayúscula, minúscula, número, etc.
)

func ValidarContrasena(c string) bool {

	// Early return (Retorno temprano): Si mide menos de 8 caracteres o tiene espacios en blanco,
	// la rechazo inmediatamente sin gastar memoria revisando letra por letra.
	if len(c) < 8 || strings.Contains(c, " ") {
		return false
	}

	// Declaro e inicializo mis 3 contadores de golpe para ahorrar líneas.
	mayusculas, numeros, especiales := 0, 0, 0

	// Recorro la contraseña caracter por caracter ('r' es el caracter actual).
	for _, r := range c {
		
		// Verifico qué tipo de caracter es y sumo al contador correspondiente.
		// Uso 'if / else if' porque un caracter no puede ser mayúscula y número al mismo tiempo.
		if unicode.IsUpper(r) {
			mayusculas++
		} else if unicode.IsDigit(r) {
			numeros++
		} else if strings.ContainsRune("!@#$%&*", r) {
			especiales++
		}
	}

	// Retorno directamente el resultado de evaluar todos los requisitos a la vez.
	// Si tiene al menos 2 mayúsculas Y al menos 1 número Y al menos 1 especial, devuelve true.
	return mayusculas >= 2 && numeros >= 1 && especiales >= 1
}

func main() {
	// Le pedimos la contraseña al usuario
	fmt.Print("Ingresa una contraseña para validar: ")

	// Vuelvo a usar bufio porque 'fmt.Scan' ignora los espacios, 
	// y mi requerimiento me exige detectar si el usuario metió un espacio para invalidarlo.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()

	// Llamo a la función y evaluo directamente el booleano que me regresa
	if ValidarContrasena(entrada) {
		fmt.Println("\nResultado: ¡Contraseña Válida! Cumple con todos los requisitos de seguridad.")
	} else {
		fmt.Println("\nResultado: Contraseña Inválida.")
		fmt.Println("Recuerda: Mínimo 8 caracteres, 2 mayúsculas, 1 número, 1 especial (!@#$%&*) y sin espacios.")
	}
}