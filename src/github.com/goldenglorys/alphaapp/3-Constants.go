package main

import (
	"fmt"
)

const (
		// _ = iota // Got generated and thrown away by compiler / Ignore first value by assigning to blank identifier
		errorSpeciallist = iota
		catSpeciallist
		dogSpeciallist
		snakeSpeciallist
)

const (
		isAdmin = 1 << iota // 1 bit-shifted 0 places -> 1
		isHeadquarters		// 1 bit-shifted 1 places -> 2
		canSeeFinancials	// 1 bit-shifted 2 places -> 4

		canSeeAfrica		// 1 bit-shifted 4 places -> 8
		canSeeAsia			// 1 bit-shifted 8 places -> 16
		canSeeEurope		// 1 bit-shifted 16 places -> 32
		canSeeNorthAmerica	// 1 bit-shifted 32 places -> 64
		canSeeSouthAmerica	// 1 bit-shifted 64 places -> 128
) // each occupy one location in a byte

func main()  {
	// Typed Constant
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
	// Untyped Constant
	const a = 42
	var b int16 = 27
	fmt.Printf("%v, %T\n", a + b, a + b)
	// Enumerated Constant
	fmt.Printf("%v, %T\n", errorSpeciallist, errorSpeciallist)
	fmt.Printf("%v, %T\n", catSpeciallist, catSpeciallist)
	fmt.Printf("%v, %T\n", dogSpeciallist, dogSpeciallist)
	fmt.Printf("%v, %T\n", snakeSpeciallist, snakeSpeciallist)
	var speciallistType int
	fmt.Printf("%v\n", speciallistType == errorSpeciallist)
	fmt.Printf("%v\n", speciallistType == catSpeciallist)

	// Bit-Shifting (to set boolean flags inside of single byte)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters & roles == isHeadquarters)	
}