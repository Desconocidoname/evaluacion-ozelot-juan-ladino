package main

import (
	"bufio" // leer lo que el usuario escriba en la terminal
	"fmt"   // imprimir en consola
	"os"    // acceder a la terminal del sistema
	"time"  // Librería nativa de Go para hacer cálculos de fechas y horas
)

func CalcularTiempo(f1, f2 string) map[string]float64 {
	
	formato := "2006-01-02 15:04"

	t1, _ := time.Parse(formato, f1)
	t2, _ := time.Parse(formato, f2)

	// h guardará las horas exactas con decimales 
	h := t2.Sub(t1).Hours()

	// Retornamos el diccionario sin envolver la matemática en 'int()'.
	return map[string]float64{
		"Años":    h / 8766,  
		"Meses":   h / 730.5, 
		"Semanas": h / 168,   
		"Días":    h / 24,    
		"Horas":   h,         
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Formato requerido: AÑO-MES-DIA HORA:MINUTO (Ejemplo: 2020-05-15 14:30)")
	
	fmt.Print("Ingresa la fecha y hora INICIAL: ")
	scanner.Scan()
	fechaInicial := scanner.Text()

	fmt.Print("Ingresa la fecha y hora FINAL: ")
	scanner.Scan()
	fechaFinal := scanner.Text()

	resultado := CalcularTiempo(fechaInicial, fechaFinal)

	fmt.Println("\n--- Diferencia de tiempo ---")
	for clave, valor := range resultado {
		fmt.Printf("%s transcurridos: %.2f\n", clave, valor)
	}
}