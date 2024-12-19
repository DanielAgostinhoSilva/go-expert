package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 10000, "Jose": 20000, "Maria": 30000}
	fmt.Println(salarios)
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 50000
	fmt.Println(salarios)
	for k, v := range salarios {
		fmt.Println(k, v)
	}
	sal := make(map[string]int)
	sal["Daniel"] = 10000
	fmt.Println(sal)
}
