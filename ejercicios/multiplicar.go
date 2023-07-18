package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error
var texto string

func Tablamultiplicar() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("ingrese un numero")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}

	}

	for i := 1; i <= 12; i++ {

		//	fmt.Printf("%d x %d = %d \n", i, num, i*num)

		texto += fmt.Sprintf("%d x %d = %d \n", i, num, i*num)

	}

	fmt.Println(texto)
	return texto
}
