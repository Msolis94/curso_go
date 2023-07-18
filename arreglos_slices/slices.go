package arreglos_slices

import (
	"fmt"
)

var table []int = []int{2, 5, 6}
var arreglo [10]int = [10]int{1, 5, 6, 8, 65, 8, 9, 9, 8, 3}

func MuestroSlice() {

	porcion := arreglo[3:]   //slice creado dede un vector, desde la posicion 3
	porcion2 := arreglo[:5]  //slice creado desde el 0 al 5
	porcion3 := arreglo[6:7] //slice creado

	fmt.Println(table)
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elemento := make([]int, 5, 20)
	fmt.Printf("Largo %d, capacidad %d \n", len(elemento), cap(elemento))

	num := make([]int, 0)

	for i := 0; i < 100; i++ {
		num = append(num, i)
	}

	fmt.Printf("Largo %d, capacidad %d", len(num), cap(num))
}
