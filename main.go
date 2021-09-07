package gocommandlineshoppingapp

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)


func main() {
	username := flag.String("username", "", "Enter your username!")
	password := flag.String("password", "", "Enter your password!")
	flag.Parse()

	loginUser(*username, *password, "users.csv")
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