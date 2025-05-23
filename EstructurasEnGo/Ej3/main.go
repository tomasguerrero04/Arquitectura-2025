package main

import (
	"fmt"
	"math/rand"
)

func main() {
	menu()

}

func play() {

	numAleatorio := rand.Intn(100)
	var numeroIngresado int

	for CantIntentos := 0; CantIntentos < 10; CantIntentos++ {
		fmt.Printf("Ingrese un numero (intentos restantes %d): ", 10-CantIntentos)
		fmt.Scanln(&numeroIngresado)
		if numeroIngresado == numAleatorio {
			fmt.Println("Ganaste")
			menu()
		} else if numeroIngresado < numAleatorio {
			fmt.Println("Incorrecto, el numero es mayor")
		} else if numeroIngresado > numAleatorio {
			fmt.Println("Incorrecto, el numero es menor")
		}
	}
	fmt.Println("Perdiste, el numero era: ", numAleatorio)
	menu()
}

func menu() {
	fmt.Println("1. Jugar")
	fmt.Println("2. Salir")
	var opcion int
	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		{
			play()
		}
	case 2:
		{
			fmt.Println("Saliendo...")
			return
		}
	default:
		{
			fmt.Println("Opcion no valida")
			menu()

		}

	}
}
