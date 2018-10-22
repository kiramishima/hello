package main

import (
	"github.com/kiramishima/hello/models"
	"github.com/kiramishima/hello/pkg"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	pkg.Message("Kira Mishima\n")
	p := pkg.SampleWithStruct("Kira Mishima\n")
	fmt.Printf("%T\n", p)
	pkg.Message(p.Name)
	langs := []string {"PHP", "Golang", "C#"}
	fmt.Println(langs)
	d := models.Developer{pkg.Person {"Kira Mishima"}, langs}
	fmt.Printf("%s\n", d)
	d2 := models.Developer{Person: pkg.Person {"Kira Mishima"}, Langs: []string {"PHP", "Golang", "C#"}}
	fmt.Println(d2)
}
