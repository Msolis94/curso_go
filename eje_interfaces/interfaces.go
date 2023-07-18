package eje_interfaces

import (
	"curso_go/interfaces"
	"fmt"
)

func HumanosRespirando(this interfaces.Humano) {
	this.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", this.Sexo())
}
