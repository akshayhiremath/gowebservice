package main

import (
	"fmt"

	"github.com/akshayhiremath/gowebservice/models"
)

func main() {
	u1 := models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Println(u1)
}
