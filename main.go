package main

import (
	"github.com/rokin04/go/models"
	"fmt"
)

func main() {
	u := models.User{
		ID: 2,
		Firstname: "rake",
		Lastname: "m",
	}
	fmt.Println(u)
}
