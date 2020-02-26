package customer

import "fmt"

//Customer ...
type Customer struct {
	Name string
}

//Greet ...
func (c Customer) Greet() {
	fmt.Println("Hello, World!")
}
