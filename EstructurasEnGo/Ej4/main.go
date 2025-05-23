package main

import (
	"fmt"
)

type Tarea struct {
	nombre      string
	completada  bool
	descripcion string
}

type ListaTareas struct {
	tareas []Tarea
}

func (lista *ListaTareas) MostrarTareas() {
	fmt.Printf("Lista de tareas:\n")
	for i, t := range lista.tareas {
		fmt.Printf("%d, %s, %s, completada = %t", i, t.nombre, t.descripcion, t.completada)

	}
}

func main() {
	var opcion int

	fmt.Print(
		"1. Agregar tarea \n",
		"2. Marcar tarea como completada \n",
		"3. Editar Tarea \n",
		"4. Eliminar Tarea 'n",
		"5. Salir \n",
		"Ingrese la opción: \n",
	)

	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		{ // Agregar tarea+
		}
	case 2:
		{ // Marcar tarea como completada
		}
	case 3:
		{ // Editar tarea
		}
	case 4:
		{ // Eliminar tarea
		}
	case 5:
		{ // Salir
			fmt.Println("Saliendo...")
			return

		}
	}
}
