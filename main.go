package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type product struct {
	id int
	name string
	price float64
	quantity int
}


func main() {
	username := flag.String("username", "", "Enter your username!")
	password := flag.String("password", "", "Enter your password!")
	flag.Parse()


	loginUser(*username, *password, "users.csv")
	listOfProducts := getProducts("products.csv")

	
}

func getProducts(productFile string) []product {
	
	productCSV, err := os.Open(*&productFile)
	if err != nil {
		fmt.Printf("Unable to read %s file", productFile)
	}

	r := csv.NewReader(productCSV)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Enable to parse %s", productFile)
	}
	items := make([]product, len(lines))
	for i, line := range lines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		price, _ := strconv.ParseFloat(line[2], 64)
		quantity, _ := strconv.Atoi(line[3])

		items[i] = product{
			id,
			name,
			price,
			quantity,
		}
	}

	return items
}



func loginUser(username string, password string, usersFile string) bool {
	csvFile, err := os.Open(usersFile)

	if err != nil {
		fmt.Printf("Unable to read %s file", usersFile)
	}

	r := csv.NewReader(csvFile)

	line, err := r.ReadAll()

	if err != nil {
		fmt.Println("Unable to parse the CSV file")
		os.Exit(1)
	}

	for _, login := range line {
		if (login[0] == username) && (login[1] == password) {
			fmt.Println("Successfully logged in!")
			return true
		}
	}
	fmt.Println("Loggin failed!")
	return false
}