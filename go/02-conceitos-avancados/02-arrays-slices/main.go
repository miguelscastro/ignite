package main

import "fmt"

func main() {
	// um slice é um array dinâmico que cresce e diminui em tamanho. No fim cada slice é um ponteiro para um array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	arr[2] = 15
	slice[0] = 123
	fmt.Println(slice)
}
