package main

import (
	"fmt"
	"math"
)

func main() {

	var lado1, lado2 float64

	fmt.Println("Ingrese el valor del lado 1:")
	fmt.Scanln(&lado1)
	fmt.Println("Ingrese el valor del lado 2:")
	fmt.Scanln(&lado2)

	area := (lado1 * lado2) / 2
	fmt.Println(" El area del triangulo es:", area)
	hipotenusa := math.Sqrt(lado1*lado1 + lado2*lado2)
	fmt.Println(" La hipotenusa del triangulo es :", hipotenusa)
	perimetro := lado1 + lado2 + hipotenusa
	fmt.Println(" El perimetro del triangulo es :", perimetro)

}
