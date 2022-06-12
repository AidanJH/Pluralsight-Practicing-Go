package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	handle string
	name   string
}

//This function can not be found publicly unless you use a Type Alias (with the = sign)
func (h Handler) RandomStuff() {

}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://twitter.com/%s\n", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

//Can use in place of a constructor
func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	//We are setting this here to prevent errors for 0 length handlers, prevents null maybe?
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		//As this method is returning an error, it should return an error if one is detected, otherwise return nothing (signifies success)
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
