package main

func suma(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a + b
}

func main() {

	println(suma(8, 4))
}
