package Teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var leyenda string
var err error

func Ingresonumero() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer numero")

	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese el 2do numero")

	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda")

	if scanner.Scan() {
		leyenda = scanner.Text()

		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println(leyenda, num1*num2)

}
