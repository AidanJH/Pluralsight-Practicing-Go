package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("James", "Anderson")
	err := p.SetTwitterHandler("@james")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("Error occured: %s\n", err.Error())
	}
	fmt.Println(p.FullName())
	fmt.Println(p.TwitterHandler().RedirectUrl())
}
