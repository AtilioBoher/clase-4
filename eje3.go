package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func show(x chacra) {
	fmt.Println("nombre: ", x.nombre)
	fmt.Println("frutas:", x.fruta)
	fmt.Println("temperatura:", x.temp)
	fmt.Println("humedad:", x.humedad)
}

func agregar(x chacra) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-2]
	if s, err := strconv.ParseFloat(text, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	// x.temp = append(x.temp, y)
}

type chacra struct { //perfil bancario
	temp    []float32
	humedad []float32
	fruta   map[string]int
	nombre  string
}

func main() {
	chacra1 := chacra{}
	chacra1.temp = append(chacra1.temp, 1, 2, 3, 4)
	chacra1.humedad = append(chacra1.humedad, 2, 3, 4, 5)
	chacra1.fruta = make(map[string]int)
	chacra1.fruta["manzana"] = 10
	chacra1.fruta["banana"] = 11
	chacra1.nombre = "fulano"

	chacra2 := chacra{}
	chacra2.temp = append(chacra2.temp, 10.1, 9, 8, 7)
	chacra2.humedad = append(chacra2.humedad, 9, 8, 7, 6)
	chacra2.fruta = make(map[string]int)
	chacra2.fruta["limón"] = 12
	chacra2.fruta["naranjas"] = 13
	chacra2.nombre = "mengano"

	show(chacra1)
	fmt.Println("")
	show(chacra2)

	agregar(chacra1)
}
