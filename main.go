package main

//go run main.go para correr
//go builder compilar
//go mod init path generar  go_mod

/*"fmt"
"runtime"
"curso_go/Teclado"*/
import (
	//"curso_go/funciones"
	//"curso_go/arreglos_slices"
	//"curso_go/mapas"
	//"curso_go/Modelos"
	//"curso_go/users"
	//m "curso_go/Modelos"
	//e "curso_go/eje_interfaces"

	//e "curso_go/defer_panic"
	//e "curso_go/gorotines"
	//"fmt"
	//e "curso_go/webserver"
	e "curso_go/middleware"
)

func main() {

	/*Variables.MostrarEnteros()
	Variables.RestoVariable()*/

	/*Estado, Texto := Variables.ConvertNoaText(1522)

	fmt.Println(Estado)
	fmt.Println(Texto)*/

	/*if os := runtime.GOOS; os == "linux." {
		fmt.Println("es linux")

	} else {
		fmt.Println("es :", runtime.GOOS)
	}*/

	/*switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("es linux")
	case "OS X":
		fmt.Println("es windows")

	default:
		fmt.Printf("%s \n", os)
	}*/

	/*Teclado.Ingresonumero()*/
	//file.LeoArchivo2()

	//funciones.Calculos()

	//funciones.LlamarClosure()
	//funciones.Exponencia(5)
	//arreglos_slices.MuestroArreglo()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapa()
	//users.AltaUsuario()

	//	pedro := new(m.Hombre)
	//e.HumanosRespirando(pedro)
	//e.VemosDefer()
	//e.EjemploPanic().
	//colocalar go antes del llamado lo convierte en asincrona
	//los chanel son para detenerlo
	/*canal1 := make(chan bool)
		go e.MiNombreLento("marlon", canal1)
		fmt.Println("estoy aqui")
	//await en go
		<-canal1*/
	//e.MiWebServer()
	e.MiMiddleware()
}
