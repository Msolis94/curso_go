package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}
func resta(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {

	fmt.Println("inicio")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(resta)(5, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

//devuelve lo que recibe
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("inicio de operacion")
		return f(a, b)
	}
}
