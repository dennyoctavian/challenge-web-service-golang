package main

import (
	"fmt"
)

func main() {

	for j := 0; j <= 10; j++ {
		if j == 0 {
			for i := 0; i < 5; i++ {
				fmt.Println("Nilai i =", i)
			}
		}
		if j == 5 {
			var arrayRussian = [7]rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'}
			for i := 0; i < len(arrayRussian); i++ {
				var sentence = fmt.Sprintf("character %U '%c' starts at byte position %d\n", arrayRussian[i], arrayRussian[i], i*2)
				fmt.Print(sentence)
			}
			continue
		}
		fmt.Println("Nilai j =", j)
	}
}
