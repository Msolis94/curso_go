package Variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariable() {

	Nombre = "pedro"
	Estado = true
	Sueldo = 1556.23
	Fecha = time.Now()

	fmt.Println("Nombre", Nombre)
	fmt.Println("Estado", Estado)
	fmt.Println("Sueldo", Sueldo)
	fmt.Println("Fecha", Fecha)

}

func ConvertNoaText(numero int) (bool, string) {

	//var Texto  string

	Texto := strconv.Itoa(numero)

	return true, Texto
}
