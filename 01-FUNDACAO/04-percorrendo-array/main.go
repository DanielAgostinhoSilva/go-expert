package main

import "fmt"

func main() {
	var meuArryar [5]int
	meuArryar[0] = 10
	meuArryar[1] = 20
	meuArryar[2] = 30
	meuArryar[3] = 40
	meuArryar[4] = 50

	fmt.Printf("O primeiro indice do array é %d e seu valor %v\n", 0, meuArryar[0])
	fmt.Printf("O ultimo indice do array é %d e seu valor %v\n", len(meuArryar)-1, meuArryar[len(meuArryar)-1])

	for i, v := range meuArryar {
		fmt.Printf("O indice %d tem o valor %d\n", i, v)
	}
}
