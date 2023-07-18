package defer_panic

import (
	//    "os"
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer msj")
	defer fmt.Println("Este es el msj final")
	fmt.Println("Este es el segundo msj ")
}

func EjemploPanic() {

	defer func() {

		reco := recover()

		if reco != nil {
			log.Fatalln("ocurrió un error que genero Panic \n", reco)

		}

	}()

	a := 1
	//se manda el panic al recover y van de las manos funciona como try catch
	if a == 1 {
		panic("Se encontro el valor 1 ")
	}
}
