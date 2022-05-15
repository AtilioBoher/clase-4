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

func readFloat() (f float64) {
	text := readString()
	f, _ = strconv.ParseFloat(text, 64)
	return f
}

func readString() (text string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	return text[:len(text)-2]
}

func agregar(x *chacra) {
	fmt.Println("Ingrese valor de temperatura: ")
	aux := readFloat()
	x.temp = append(x.temp, aux)
	fmt.Println("Ingrese valor de humedad: ")
	aux = readFloat()
	x.humedad = append(x.humedad, aux)
}

func promedioTemp(x chacra) (f float64) {
	var sum float64
	for _, v := range x.temp {
		sum += v
	}
	return sum / float64(len(x.temp))
}

func promedioHum(x chacra) (f float64) {
	var sum float64
	for _, v := range x.humedad {
		sum += v
	}
	return sum / float64(len(x.humedad))
}

func datosFruta(x chacra, nombre string) {
	fruta, existe := x.fruta[nombre]
	if existe == true {
		fmt.Printf("para la futa \"%s\" existen %d unidades", nombre, fruta)
	} else {
		fmt.Printf("la fruta \"%s\" no está en el registro", nombre)
	}
}

type chacra struct { //perfil bancario
	temp    []float64
	humedad []float64
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

	fmt.Println("chacra1")
	show(chacra1)
	fmt.Println("chacra2")
	show(chacra2)

	agregar(&chacra1)
	fmt.Println("chacra1 con valores agregados")
	show(chacra1)

	fmt.Println("el promedio de temperatura es %2.f", promedioTemp(chacra1))
	fmt.Println("el promedio de humedad es %2.f", promedioHum(chacra1))
	datosFruta(chacra2, "limón")

}
