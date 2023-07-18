package funciones

import "fmt"

func tabla(valor int) func() int {
	num := valor
	secuencias := 0

	return func() int {
		secuencias++
		return num * secuencias
	}
}

func LlamarClosure() {

	tabladel := 2
	MiTabla := tabla(tabladel)
	for i := 1; i <= 12; i++ {
		fmt.Println(MiTabla())
	}

}
