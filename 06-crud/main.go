package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Tienda struct {
	db   map[int]Producto
	auth bool
}

// Leo el archivo .env manualmente sin usar librerías de terceros.
// Si el archivo no existe o hay un error, las variables regresan vacías y el login fallará por seguridad.
func cargarCredenciales() (usr, pass string) {
	archivo, err := os.Open("example.env")
	if err != nil { return }
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		p := strings.Split(scanner.Text(), "=")
		if len(p) == 2 {
			if p[0] == "USUARIO" { usr = strings.TrimSpace(p[1]) }
			if p[0] == "PASSWORD" { pass = strings.TrimSpace(p[1]) }
		}
	}
	return
}

func (t *Tienda) Login(usr, pass string) {
	envUsr, envPass := cargarCredenciales()
	// Verifico que coincidan y que no estén vacías (por si falta el archivo .env)
	if usr == envUsr && pass == envPass && usr != "" {
		t.auth = true
		fmt.Println("Autenticación exitosa.")
	} else {
		fmt.Println("Error: Credenciales incorrectas o falta el archivo .env.")
	}
}

// Validador centralizado.
func (t *Tienda) checkAuth() bool {
	if !t.auth { fmt.Println("Error: requiere iniciar sesión."); return false }
	return true
}

func (t *Tienda) Listar() {
	if !t.checkAuth() { return }
	fmt.Println("\n--- Inventario ---")
	for id, p := range t.db {
		fmt.Printf("ID: %d | %s | $%.2f | Stock: %d\n", id, p.Nombre, p.Precio, p.Cantidad)
	}
}

// Un mapa permite agregar un elemento nuevo o sobrescribir uno existente (editar) 
// con la misma instrucción. Así evito repetir código.
func (t *Tienda) Guardar(id int, p Producto) {
	if !t.checkAuth() { return }
	t.db[id] = p
	fmt.Println("Producto guardado.")
	t.Listar()
}

func (t *Tienda) Eliminar(id int) {
	if !t.checkAuth() { return }
	if _, ok := t.db[id]; !ok { fmt.Println("Error: ID no existe."); return }
	delete(t.db, id)
	fmt.Println("Producto eliminado.")
	t.Listar()
}

// Función auxiliar para leer texto en terminal sin que los espacios rompan el input.
func pedirDato(s *bufio.Scanner, msg string) string {
	fmt.Print(msg)
	s.Scan()
	return s.Text()
}

func main() {
	// Inicializo el mapa directamente al crear la estructura.
	app := &Tienda{db: make(map[int]Producto)}
	scanner := bufio.NewScanner(os.Stdin)

	usr := pedirDato(scanner, "Usuario: ")
	pass := pedirDato(scanner, "Contraseña: ")
	app.Login(usr, pass)

	if !app.auth { return }

	// Menú infinito usando switch/case
	for {
		fmt.Println("\n1. Agregar/Editar | 2. Eliminar | 3. Listar | 4. Salir")
		switch pedirDato(scanner, "Opción: ") {
		
		case "1":
			id, _ := strconv.Atoi(pedirDato(scanner, "ID: "))
			nombre := pedirDato(scanner, "Nombre: ")
			pre, _ := strconv.ParseFloat(pedirDato(scanner, "Precio: "), 64)
			can, _ := strconv.Atoi(pedirDato(scanner, "Stock: "))
			
			// Paso todo directamente a la función Guardar.
			app.Guardar(id, Producto{nombre, pre, can})
			
		case "2":
			id, _ := strconv.Atoi(pedirDato(scanner, "ID a eliminar: "))
			app.Eliminar(id)
			
		case "3":
			app.Listar()
			
		case "4":
			return 
			
		default:
			fmt.Println("Opción inválida.")
		}
	}
}