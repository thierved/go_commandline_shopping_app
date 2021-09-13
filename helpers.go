package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"
	// "strings"
)

func readCSV(file string) [][]string {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to read %s file", file)
	}
	r := csv.NewReader(csvFile)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse the CSV file")
		os.Exit(1)
	}
	return lines
}

func selectItems(message string) []string {
	var input string 
	fmt.Println(message)
	fmt.Scanf("%s\n", &input)
	return strings.Split(input, "-")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}