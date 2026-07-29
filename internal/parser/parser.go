package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

//lê arquivos CSV e os coloca num struct
func ParseCSV(fileArg string) ([]Record, error){
	var pessoas []Record

	file, err := os.Open(fileArg)
	if err != nil{
		return nil, fmt.Errorf("%w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil{
		return nil, fmt.Errorf("erro ao ler CSV: %w", err)
	}
	for {
		record, err := reader.Read()
		if err == io.EOF{
			break
		}
		if err != nil{
			return nil, fmt.Errorf("erro ao ler CSV: %w", err)
		}
		if len(record) != 3{
			return nil, fmt.Errorf("erro na formatacao do CSV: ID,Nome,email")
		}
		id, err := strconv.Atoi(record[0])
		if err != nil{
			return nil, fmt.Errorf("erro na conversão do ID %v: %w", record[0], err)
		}
		pessoas = append(pessoas, Record{
			ID: id,
			Name: record[1],
			Email: record[2],
		})
	}

	return pessoas, nil
}