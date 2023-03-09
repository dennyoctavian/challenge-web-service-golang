package main

import (
	"fmt"
)

func main() {
	i := 21
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Println("%")
	j := true
	fmt.Println(j)
	fmt.Println()
	fmt.Printf("%b \n", i)
	russian_char := "\u042F"
	fmt.Printf("%q\n", russian_char)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)

	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U  \n", 'Я')
	k := 123.456

	fmt.Println()
	fmt.Printf("%f  \n", k)
	fmt.Printf("%E  \n", k)

}
