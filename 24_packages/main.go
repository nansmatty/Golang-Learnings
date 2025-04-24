package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nansmatty/testing/auth"
	"github.com/nansmatty/testing/user"
)

func main() {

	auth.LoginWithCredentials("Narayan", "123456789")
	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Name:     "Narayan Maity",
		Email:    "nansmatty97@gmail.com",
		Password: "123456789",
	}

	fmt.Println(user.Name)
	fmt.Println(user.Email)

	color.Red(user.Email)
	color.Green(user.Email)

}
