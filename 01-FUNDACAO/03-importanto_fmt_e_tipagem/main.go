package main

import "fmt"

type ID int
type NOME string

var (
	a ID   = 1234
	b NOME = "teste"
	c bool = false
)

func main() {
	fmt.Printf("O tipo de A é %T e seu valor é %v\n", a, a)
	fmt.Printf("O tipo de B é %T e seu valor é %v\n", b, b)
	fmt.Printf("O tipo de C é %T e seu valor é %v\n", c, c)
}
