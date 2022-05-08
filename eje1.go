package main

import "fmt"

func promedio(x []int) float32 {
	sum := 0
	for _, v := range x {
		sum += v
	}
	sumFloat := float32(sum)
	aux := float32(len(x))
	return sumFloat / aux
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("todos los números => ", numeros)

	fmt.Print("\nlos siguientes números son impares:")
	for _, v := range numeros {
		if v%2 != 0 {
			fmt.Print("\n", v)
		}
	}
	newSlice := numeros[3:7]
	fmt.Print("\nnumeros[3:7] => ", newSlice)

	fmt.Print("\npromedio => ", promedio(numeros))

}
