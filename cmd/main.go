package main

import (
	"fmt"
	"log"
	"os"
	"github.com/ArthurBryan166/Estudos/estudo-de-cli/leitor-arquivocsv/internal/parser"
)

func main(){
	args := os.Args
	if len(args) != 2{
		log.Fatalln("uso: app <arquivo.csv>")
	}

	cadastros, err := parser.ParseCSV(args[1])
	if err != nil{
		fmt.Println(err)
		return
	}

	for _, v := range cadastros{
		fmt.Printf("ID: %v\nNome: %v\nEmail: %v\n\n", v.ID, v.Name, v.Email)
	}
}