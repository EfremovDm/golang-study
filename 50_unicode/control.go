package main

import "fmt"

/**
Unicode
 */
func main() {

	var symbol rune = 'a'
	var autoSymbol = 'a'
	unicodeSymbolNumber := '\u2318'

	println(symbol, autoSymbol, unicodeSymbolNumber)
	fmt.Println(symbol, autoSymbol, unicodeSymbolNumber)

	str1 := "Привет, Мир!"
	fmt.Println("ru: ", str1, len(str1))
	for index, runeValue := range str1 {
		fmt.Println("%#U as position %d\n", runeValue, index)
	}
	println(str1[2])
}