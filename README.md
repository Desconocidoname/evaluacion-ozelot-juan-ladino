# Evaluación Diagnóstica - Ozelot Technologies

**Candidato:** Juan Daniel Ladino Hernandez

El presente repositorio contiene la solución a los seis ejercicios algorítmicos requeridos para la evaluación diagnóstica de prácticas profesionales en Ozelot Technologies. 

Los ejercicios fueron desarrollados íntegramente en el lenguaje de programación **Go (Golang)**.

---

## Requisitos Previos e Instalación

Para compilar y ejecutar los algoritmos de este repositorio, es necesario contar con el entorno de Go instalado en el sistema.

**Instalación en entornos Arch Linux:**
\`\`\`bash
sudo pacman -S go
\`\`\`

**Instalación en otras plataformas (Windows / macOS):**
Se recomienda descargar el instalador oficial directamente desde la documentación de Go: [go.dev/doc/install](https://go.dev/doc/install).

Para verificar que la instalación se configuró correctamente, ejecute el siguiente comando en la terminal:
\`\`\`bash
go version
\`\`\`

---

## Instrucciones de Ejecución

Cada algoritmo se encuentra encapsulado en su propio directorio para asegurar su independencia estructural. Para evaluar las soluciones, abra la terminal en la raíz de este repositorio y ejecute los comandos correspondientes a cada módulo de forma secuencial:

### 1. Validador de Palíndromos (Básico)
```bash
cd 01-palindromo
go run main.go
```

### 2. Cálculo de Tiempo Transcurrido (Básico)
```bash
cd ../02-tiempo
go run main.go
```

### 3. Validador de Contraseñas (Intermedio)
```bash
cd ../03-password
go run main.go
```

### 4. Adivina el Número (Intermedio)
```bash
cd ../04-adivina
go run main.go
```

### 5. Sucesión de Fibonacci (Avanzado)
```bash
cd ../05-fibonacci
go run main.go
```

### 6. CRUD de Gestión de Productos con Autenticación (Avanzado)
**Nota de Configuración Importante:**
Este módulo implementa un lector nativo de variables de entorno por motivos de seguridad. Antes de inicializar el programa, es importante verificar la existencia del archivo denominado `example.env` dentro del directorio `06-crud`.

**Ejecución del sistema:**
```bash
cd ../06-crud
go run main.go
```