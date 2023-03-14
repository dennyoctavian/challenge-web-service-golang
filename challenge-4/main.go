package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	users := []map[string]interface{}{
		{"No": 1, "Nama": "Denny Octavian", "Alamat": "Karang Anyar", "Pekerjaan": "Mobile Developer", "Alasan": "Ingin Belajar TechStack Baru"},
		{"No": 2, "Nama": "Budi Gunawan", "Alamat": "Ancol", "Pekerjaan": "Mobile Developer", "Alasan": "Ingin Belajar TechStack Baru"},
		{"No": 3, "Nama": "Asep", "Alamat": "Pademangan", "Pekerjaan": "Mobile Developer", "Alasan": "Ingin Belajar TechStack Baru"},
		{"No": 4, "Nama": "Joko", "Alamat": "Senen", "Pekerjaan": "Mobile Developer", "Alasan": "Ingin Belajar TechStack Baru"},
		{"No": 5, "Nama": "Eko", "Alamat": "Pluit", "Pekerjaan": "Mobile Developer", "Alasan": "Ingin Belajar TechStack Baru"},
		{"No": 6, "Nama": "Juki", "Alamat": "PIK", "Pekerjaan": "Mobile Developer", "Alasan": "Ingin Belajar TechStack Baru"},
		{"No": 7, "Nama": "Agung", "Alamat": "Cengkareng", "Pekerjaan": "Mobile Developer", "Alasan": "Ingin Belajar TechStack Baru"},
	}

	arguments := os.Args

	if len(arguments) == 2 {
		intVar, err := strconv.ParseInt(arguments[1], 10, 64)

		var result map[string]interface{}

		if err != nil {
			fmt.Println("Input yang anda masukan salah")
		} else {
			for _, v := range users {
				if int64(v["No"].(int)) == intVar {
					fmt.Println("ada yang match")
					result = v
				}
			}
			if result == nil {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println(result)
			}

		}

	} else {
		fmt.Println("Input yang anda masukan salah")
	}

}
