package file

import (
	"curso_go/ejercicios"
	"fmt"

	"bufio"
	"io/ioutil"
	"os"
	//"ioutil"
)

var filename string = "./file/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.Tablamultiplicar()
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear el Archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Tablamultiplicar()

	if !Append(filename, texto) {
		fmt.Println("error al concatenar el contenido")

	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {

		fmt.Println("Error durante" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {

		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

//funcion para leer archivo

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(filename)

	if err != nil {

		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo))
}

func LeoArchivo2() {
	archivo, err := os.Open(filename)

	if err != nil {

		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
