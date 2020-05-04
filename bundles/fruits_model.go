package bundles

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Fruit struct {
	FruitName string  `json:"fruit" validate:"required"`
	Calories  int64   `json:"calories" validate:"required,gte=0,lte=400"`
	Price     float64 `json:"price" validate:"gte=0,lte=100"`
}



func LoadCSV(path string)[]Fruit {
	fmt.Println("LoadCSV#############")
	csvfile, err := os.Open(path)
	r := csv.NewReader(csvfile)
	var f []Fruit = nil
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Unable to parse csv file", err)
		}
		cal, err := strconv.ParseInt(record[1], 10, 64)
		price, err := strconv.ParseFloat(record[2], 2)
		f = append(f, Fruit{
			FruitName: record[0],
			Calories:  cal,
			Price:     price,
		})

	}

	return f
}

func WriteCSV(path string, f []Fruit) {
	csvFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvFile.Truncate(0)
	csvFile.Seek(0,0)
	csvwriter := csv.NewWriter(csvFile)

	for _, fruit := range f {
		var line = []string{fruit.FruitName, strconv.Itoa(int(fruit.Calories)), fmt.Sprintf("%f", fruit.Price)}
		_ = csvwriter.Write(line)
	}
	csvwriter.Flush()
	csvFile.Close()
}
