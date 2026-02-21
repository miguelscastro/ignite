package main

import "fmt"

type Foo struct {
	Bar string
	Baz float64
	Quz []string
}

// gopls preenche automaticamente usando o fillStruct via ctrl + .
func FillStruct() {
	f := Foo{
		Bar: "",
		Baz: 0,
		Quz: []string{},
	}
	fmt.Println(f)
}
