package main

import "fmt"

func loginUser(username string, password string, usersFile string) bool {
	for _, login := range readCSV(usersFile) {
		if (login[0] == username) && (login[1] == password) {
			fmt.Println("Successfully logged in!")
			return true
		}
	}
	fmt.Println("Loggin failed!")
	return false
}