package main

type ID int
type NOME string

var (
	ID_PADRAO   ID   = 0
	NOME_PADRAO NOME = "TESTE"
)

func main() {
	println(ID_PADRAO)
	println(NOME_PADRAO)
}
