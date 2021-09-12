package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "", "Enter your username!")
	password := flag.String("password", "", "Enter your password!")
	flag.Parse()


	if loginUser(*username, *password, "users.csv") {
		listOfProducts := getProducts("products.csv")
		displayProducts(listOfProducts)
		total := makeInvoice(selectProducts(selectItems("Select the items of your choice separeted by '-'"), listOfProducts))
		fmt.Printf("total coast: $%.2f\n",total)
	}
}


