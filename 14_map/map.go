package main

import "fmt" // пакет стандартной библиотеки для форматирования ввода

/**
 * Мапы, они же хеши, они же ассоциативные массивы
 * Мапы присвваиваются по ссылке
 * Порядок ключей не детерминировал (случайный порядок)
 */
func main() {

	// инициализация пустой мапы
	var mm map[string]string //[тип ключа]тип данных
	fmt.Println("uninitialized map", mm)

	var mm2 map[string] map[string]string //массив массивов
	fmt.Println("uninitialized mm2", mm2)

	// mm["test"] = "ok" // так нельзя

	// полная инициализация мапы
	var mm3 map[string]string = map[string]string{}
	fmt.Println(mm3)

	mm4 := map[string]string{}

	// добавление элемента в мапу полной инициализации
	mm4["test"] = "ok"
	fmt.Println(mm4)

	// полная инициализация мапы через make
	var mm5 = make(map[string]string)
	mm5["firstName"] = "Vasiliy"
	fmt.Println(mm5)

	// получение значения
	firstName := mm5["firstName"]
	fmt.Println("firstName", firstName)

	// обращение к несуществующему элементу - отдается значение по умолчанию
	lastName := mm5["lastName"];
	fmt.Println("lastName", lastName, len(lastName))

	// проверка на то что значение есть
	lastName, ok := mm5["lastNamr"]
	fmt.Println("lastName", lastName, "exists: ", ok)

	// получение только признака существования
	_, exist := mm5["123123"]
	fmt.Println("exists", exist)

	// удаление значения
	_, exist3 := mm5["firstName"]
	fmt.Println("exists", exist3)
	delete(mm5, "firstName")
	_, exist4 := mm5["firstName"]
	fmt.Println("exists", exist4)
}
