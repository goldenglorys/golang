package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// variable declzration
	var foo int 		// declare a variable and initialised later
	var foo2 int = 42	// declare and initialised at same time
	foo3 := 42			// short-hand version of declaration and initialisation and compiler to decide the type base on given value
	var foo4 string
	foo4 = strconv.Itoa(foo3)

	// Variable visibility
	// - lower case first letter is for pkg scope (all src can access)
	// - upper case first letter is to be export and used by any
	// - scope for a specific block without other block or pkg visibility

	// Naming variable
	// - Pascal or camelCase (capitalize acronyms URL-HTTP)
	// - longer name for longer lives

	fmt.Printf("%v, %T\n", foo, foo)
	fmt.Printf("%v %T\n", foo2, foo2)
	fmt.Printf("%v %T\n", foo3, foo3)
	fmt.Printf("%v %T\n", foo4, foo4)
	fmt.Println("Hello Go!")
}