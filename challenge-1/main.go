package main

import "fmt"

func main() {
	i := 21
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	// TODO print %
	fmt.Println("%")
	// TODO print true
	j := true
	fmt.Println(j)
	// 	Assignment 1:
	// menampilkan nilai base2 dari i : 10101
	fmt.Printf("Nilai base2 dari i adalah %b \n", i)
	// menampilkan nilai base 10 dari i : 21
	fmt.Printf("Nilai base10 dari i adalah %d \n", i)
	// menampilkan nilai base 8 dari i :25
	fmt.Printf("Nilai base8 dari i adalah %o without prefix \n", i)
	fmt.Printf("Nilai base8 dari i adalah %O with prefix \n", i)
	// menampilkan nilai base 16 : f --> base 16 dari f itu angka berapa?
	fmt.Printf("Nilai base16 dari f adalah %x with lower-case letters for a-f \n", "f")
	fmt.Printf("Nilai base16 dari f adalah %X with upper-case letters for A-F \n", "f")
	// menampilkan nilai base 16 : F --> base 16 dari f itu angka berapa?
	fmt.Printf("Nilai base16 dari F adalah %x with lower-case letters for a-f \n", "F")
	fmt.Printf("Nilai base16 dari F adalah %X with upper-case letters for A-F \n", "F")
	// menampilkan unicode karakter Я (bahasa rusia "ya") : U+042F
	fmt.Printf("Nilai UNICODE dari Я adalah %U  \n", 'Я')
	// var k float64 = 123.456;
	k := 123.456

	// menampilkan float dari k : 123.456000
	fmt.Printf("Nilai k adalah %f  \n", k)
	// menampilkan float scientific dari k: 1.234560E+02
	fmt.Printf("Nilai k dalam  float scientific adalah %E  \n", k)

}
