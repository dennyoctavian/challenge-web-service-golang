package main

import "fmt"

func main() {
	Looping("selamat malam")
}

func Looping(word string) {
	var countChar = map[string]int{}
	for _, v := range word {
		fmt.Println(string(v))
		value, exist := countChar[string(v)]

		if exist {
			countChar[string(v)] = value + 1
		} else {
			countChar[string(v)] = 1
		}
	}
	fmt.Println(countChar)
}
