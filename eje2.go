package main

import "fmt"

func show(x perBanc) {
	fmt.Println("DNI: ", x.Dni)
	fmt.Println("Nombre y Apellido: ", x.NombyApell)
	fmt.Println("Dirección: ", x.Direccion)
	fmt.Println("Historial crediticio: ", x.Nota)
	fmt.Println("Número de cuenta: ", x.numCuenta)
}

type perBanc struct { //perfil bancario
	Dni        string
	NombyApell string
	Direccion  string
	Nota       string // nota de historial crediticio
	numCuenta  int    // número de cuenta
}

func main() {

	perfil1 := perBanc{"somedni1", "some name and last name 1", "dirección1", "nota1", 100}
	perfil2 := perBanc{"somedni2", "some name and last name 2", "dirección2", "nota2", 200}

	show(perfil1)
	fmt.Print("\n")
	show(perfil2)
}
