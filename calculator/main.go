package main

import "fmt"

func main() {
	const i = 1
	for i == 1 {
		var islem string
		var sayi1, sayi2 int

		fmt.Println("sayı1 işlem sayı2?")
		fmt.Scan(&sayi1,&islem, &sayi2)

		switch islem {
		case "+":
			fmt.Println(sayi1 + sayi2)
		case "-":
			fmt.Println(sayi1 - sayi2)
		case "/":
			fmt.Println(sayi1 / sayi2)
		case "*":
			fmt.Println(sayi1 * sayi2)
		case "x":
			return
		}
	}
}
