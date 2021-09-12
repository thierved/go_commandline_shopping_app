package main

import (
	"fmt"
	"strconv"
)


type product struct {
	id int
	name string
	price float64
	quantity int
}

func displayProducts(items []product) {
	fmt.Println("============================================")
	fmt.Printf("%-3s %-20s %-10s %-8s\n", "id", "name", "price", "quantity")
	fmt.Println("============================================")
	for _, item := range items {
		fmt.Printf("%-3d %-20s $%-9.1f %-8d\n", item.id, item.name, item.price, item.quantity)
	}
}

func getProducts(productFile string) []product {
	lines := readCSV(productFile)
	items := make([]product, len(lines))
	for i, line := range lines {
		item := product{}
		item.id, _ = strconv.Atoi(line[0])
		item.name = line[1]
		item.price, _ = strconv.ParseFloat(line[2], 64)
		item.quantity, _ = strconv.Atoi(line[3])
		
		items[i] = item
	}

	return items
}

func selectProducts(selectedItems []string, productList []product) {
	pickedItems := []product{}
	for _, i := range selectedItems {
		for _, j := range productList {
			if i, _ := strconv.Atoi(i); i == j.id {
				pickedItems = append(pickedItems, j)
			}
		} 
	}	
	displayProducts(pickedItems)
}