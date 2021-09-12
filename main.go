package main

import (
	"flag"
)

func main() {
	username := flag.String("username", "", "Enter your username!")
	password := flag.String("password", "", "Enter your password!")
	flag.Parse()


	if loginUser(*username, *password, "users.csv") {
		listOfProducts := getProducts("products.csv")
		displayProducts(listOfProducts)
		selectProducts(selectItems(), listOfProducts)
	}
}


