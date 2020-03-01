package admin

import "fmt"

//Admin ...
type Admin struct{}

//SayHello ...
func (a Admin) Greet() {
	fmt.Println("Hello, I am an admin and don't you forget it.")
}
