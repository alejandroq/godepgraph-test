package main

import (
	"github.com/alejandroq/godepgraph-commithook/internal/admin"
	"github.com/alejandroq/godepgraph-commithook/internal/customer"
)

func main() {
	c := customer.Customer{}
	c.Greet()

	a := admin.Admin{}
	a.Greet()
}
