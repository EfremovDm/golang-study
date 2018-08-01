package main

import "fmt"

func main() {

	var aarr [3]int
	fmt.Println("Массив из 3-х элементов", aarr)

	arr2 := [...]int{1, 2, 3}
	fmt.Println("Длина при инициализации", arr2, " длина ", len(arr2))

	arr2[1] = 12
	fmt.Println("Измененный массив", arr2, " длина ", len(arr2))

	var arr3 [3][3]int
	arr3[1][1] = 1
	fmt.Println("Матрица", arr3)


}
