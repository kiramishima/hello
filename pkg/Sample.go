package pkg

// Person Struct de ejemplo para el modulo
type Person struct {
	Name string
}

// SampleWithStruct Ejemplo de una funcion que retorna un Struct
func SampleWithStruct(name string) Person {
	p := Person{name}
	return p
}
