package variables

import (
	"fmt"
	"strconv"
	"time"
)

var nombre string
var estado bool
var sueldo float32
var fecha time.Time

func RestoVariables() {

	nombre = "Pedro"
	estado = true
	sueldo = 2375
	fecha = time.Now()

	fmt.Println("Nombre: ", nombre)
	fmt.Println("Estado: ", estado)
	fmt.Println("Sueldo: ", sueldo)
	fmt.Println("Fecha: ", fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	stringNumero := strconv.Itoa(numero)
	fmt.Printf("El número %d convertido a string es: %s\n", numero, stringNumero)
	return true, stringNumero
}
