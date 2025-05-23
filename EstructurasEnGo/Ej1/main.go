package main

import "fmt"

const (
	MensajePrivado = "Hola desde el paquete ejemplo"
	MensajePublico = "Hola desde el paquete ejemplo"
)

const Pi float32 = 3.14
const euler float32 = 2.71828

const x = 5
const y = "hola"

func main() {
	/*
		fmt.Println(MensajePrivado)
		fmt.Println(x, y) */

	var (
		defaultInt    int
		defaultString string = "a"
		defaultFloat  float32
		defaultbool   bool
	)

	fmt.Println(defaultInt, defaultString, defaultFloat, defaultbool)

}
