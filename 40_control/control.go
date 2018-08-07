package main

/**
Управляющие конструкции
 */
func main() {

	// операторы ветвления
	a := true
	if a { // только булево значение, никаких цифр, строк и т.п.
		println("a is true")
	}

	b := 1
	c := 22
	if (b == 1 && a) || c != 22 {
		println("Неявное преобразование ( if b ) не работает")
	}

	// сложный if
	mm := map[string]string{
		"firstName": "Vasiliy",
		"lastName": "Romanov",
	}
	if firstName, ok := mm["firstName"]; ok {
		println("firstName существует, = ", firstName)
	} else {
		println("firstName не существует")
	}

	// циклы
	for  { // в go только 1 операто цикла
		println("Бесконечный цикл, аналог while (true)")
		break
	}


	s1 := []int {3,4,5,6,7,8}
	value := 0
	idx := 0

	// while
	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = s1[idx]
		idx++
		println("while, idx:", idx, " value:", value)
	}

	// for i
	for i := 0; i < len(s1); i++ {
		println("c-style loop", i, s1[i] )
	}

	// range
	for idx, val := range s1 {
		println("range slice by index", idx, val)
	}

	// операции по map
	for idx, val := range mm {
		println("range map by index", idx, val)
	}

	switch mm["firstName"] {
	case "Vasiliy", "Evgeny":
		println("switch - name is Vasiliy")
		// в отличие от других языков нет перехода на другой вариант по умолчанию
		fallthrough // переходим в следующий вариант специально
	case "Petr", "Eva":
		println("switch - name is Petr")
	default:
		println("switch - default")
	}

	// замена if else
	switch {
	case mm["firstName"] == "Vasiliy":
		println("firstName Vasiliy")
	case mm["lastName"] == "Romanov":
		println("lastName Romanov")
	}
}