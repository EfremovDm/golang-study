package main

func main() {
	//fmt.Println("Hello World!")

	// чиста в golang
	var i int = 10
	var autoint = -10
	var bigint int64 = 1<<32 - 1
	var usignedInt uint =100500
	var usignedBigInt uint64 = 1 << 64 - 1

	println("integers", i, autoint, bigint, usignedInt, usignedBigInt)
	//return

	// строки

	var hello string = "Hello\n\t"
	var world = "World"

	println(hello, world)
	//return

	// бинарные данные
	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)
	//return

	// краткое объявление
	//live := 42 // новая переменная













}
