package main

import "fmt"

type user struct {
	name string
	username string
	password string
}

func loginUser() (user, bool) {
	fmt.Println("Enter your username: ")
	var usrn string
	fmt.Scanf("%s\n", &usrn)
	fmt.Println("Enter your password: ")
	var psw string
	fmt.Scanf("%s\n", &psw)
	for _, login := range readCSV("users.csv") {		
		if (login[1] == usrn) && (login[2] == psw) {
			fmt.Println("Successfully logged in!")			
			return user{login[0], login[1], login[2]},true
		}
	}
	fmt.Println("Loggin failed!")
	return user{}, false
}