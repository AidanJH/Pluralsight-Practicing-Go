package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("James", "Anderson", organization.NewSocialSecurityNumber("125970-13"))
	err := p.SetTwitterHandler("@james")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("Error occured: %s\n", err.Error())
	}
	fmt.Println(p.FullName())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	fmt.Println(p.ID())

}
