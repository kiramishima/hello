package models

import (
	"github.com/kiramishima/hello/pkg"
)

// Developer Struct Ejemplo para jugar con el feature de modulos
type Developer struct {
	pkg.Person
	Langs []string
}
