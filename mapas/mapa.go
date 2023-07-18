package mapas

import "fmt"

func MostrarMapa() {
	//declarion por make
	paises := make(map[string]string)

	paises["mexico"] = "D.F"
	paises["Argentina"] = "Buenos aires"

	//fmt.Println(paises)
	//fmt.Println(paises["mexico"])
	//declaracion directa map
	campeonato := map[string]int{
		"barcelona": 39,
		"madrid":    50,
		"juventu":   20,
	}

	//fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s , tiene un puntaje de  %d \n", equipo, puntaje)
	}
	delete(campeonato, "juventu")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["barcelona"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)

}
