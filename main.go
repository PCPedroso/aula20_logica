package main

import "fmt"

func main() {
	var vlaor1, vlaor2, soma, media float32

	fmt.Println("Entre com dois valores: ")
	fmt.Scanf("%g", &vlaor1)
	fmt.Scanf("%g", &vlaor2)

	soma = vlaor1 + vlaor2

	media = soma / 2

	fmt.Printf("Média dos vlores %v e %v é: %v", vlaor1, vlaor2, media)

}
