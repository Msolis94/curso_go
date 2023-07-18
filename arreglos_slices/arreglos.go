package arreglos_slices

import (
	"fmt"
)

var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MuestroArreglo() {
	tabla[7] = 33
	tabla[2] = 5
	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[1][2] = 50

	fmt.Println(matriz)

}
