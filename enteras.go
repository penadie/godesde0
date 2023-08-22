package variables

import "fmt"

// funcion publica, siempre estan con mayuscula en la primera letra,
// para ser accedida de distintas partes deficincion de golanga
// el corchete debe estar en la primera fila, si lo bajas, da error
func MuestroEnteros() { // este
	//declarar variables de distintas maneras declarativa, y asignacion
	// declarada
	var intcomun int // error si no estamos usando la viable, go te lo dice
	// asignacion
	intde32 := int32(10)
	intde64 := int64(100)
	// nombre de var intde32
	// := asigna valor
	// funcion de intde32 que en resumen hace que un dato sea del tipo int de 32 bits
	// 10 es el dato que estamos fijando como tipo de int 32

	// Go puede ser muy especifico con los tipos de datos que utiliza,
	// esto tiene multiples funciones, una de ellas es para tener mejor
	// rendimiento, y no se me ocurren mas, pero debe ser amplio el tema

	fmt.Print("intcomun = ", intcomun)
	fmt.Print("intde32 = ", intde32)
	fmt.Print("intde64 = ", intde64)
	// Las variable cuando se definen con el valor minimo de ese tipo de dato
	// no hay quj exportarlo, se exporta automaticamente
	// al tener la funcion una letra mayuscula

}
