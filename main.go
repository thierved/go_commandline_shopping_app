package main

import (
	"fmt"
)

func main() {
	client, loggedIn := loginUser()
	if loggedIn {
		listOfProducts := getProducts("products.csv")
		displayProducts(listOfProducts)
		total := makeInvoice(selectProducts(selectItems("Select the items of your choice separeted by '-'"), listOfProducts), client)
		fmt.Printf("total coast: $%.2f\n",total)
	}
}


