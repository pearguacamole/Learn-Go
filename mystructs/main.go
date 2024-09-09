package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	kushagra := User{"Kushagra", "kushmail@mail.in", false, 21}
	fmt.Println(kushagra)
	fmt.Printf("Ksuhagra details are %+v\n", kushagra)
	fmt.Printf("The user name is %v and email is %v\n", kushagra.Name, kushagra.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
