package main

// глобальные переменные уровня пакета
var (
	m0 int = 12
	m2 = "string"
	m3 = 23
)

func main() {
	//fmt.Println("Hello World!")

	// чиста в golang
	var i int = 10
	var autoint = -10
	var bigint int64 = 1<<32 - 1
	var usignedInt uint =100500
	var usignedBigInt uint64 = 1 << 64 - 1

	println("integers", i, autoint, bigint, usignedInt, usignedBigInt)

	// строки
	var hello string = "Hello\n\t"
	var world = "World"

	println(hello, world)

	// бинарные данные
	var rawBinary byte = '\x27' // одинарные кавычки указываются для конкретного символа
	println("rawBinary", rawBinary)

	/**
		краткое объявление
	 */
	meaningOfLive := 42 // работает только для новых переменных
	println("meaningOfLive", meaningOfLive)

	/**
		приведение типов
	 */
	p := 3.14
	println("float to int conversion", int(p))
	println("int to string conversion", string(48)) // в результате получается номер символа, а не строка
	println("summ", int(p) + 2)

	/**
		комплексные числа о_О О_о О_О
	 */
	z := 2 + 3i
	println("complex number", z)

	/**
		операции со строками
	 */
	s1 := "Vasiliy"
	s2 := "Romanov"
	fullName := s1 + " " + s2
	escaping := `Hello

World!!!`;
	println("fullName", fullName, len(fullName))
	println("escaping", escaping)

	/**
		значение по умолчанию всегда есть во всех переменных в go, они не null
	*/
	var defaultInt int
	var defaultFloat float32
	var deafaultString string
	var de3faultBool bool
	println("defaultValues", defaultInt, defaultFloat, deafaultString, de3faultBool)

	/**
		несколько переменных
	 */
	var v1, v2 string = "v1", "v2"
	println("несколько переменных", v1,v2)
	println("несколько глобальных переменных", m0, m2, m3)
}
