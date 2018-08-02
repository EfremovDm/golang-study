package main

import "fmt" // пакет стандартной библиотеки для форматирования ввода

/**
 * Слайсы (срезы) - массивы без длины, реализованы через массивы внутри себя
 */
func main() {

	var sl1 []int

	fmt.Println("Пустой срез", sl1)
	//fmt.Println("Пустой срез", sl1[4]) // здесь будет runtime-ошибка

	sl1 = append(sl1, 100) // добавление элемента в слайс
	fmt.Println("Уже не пустой срез", sl1)

	fmt.Println("Длина внутреннего массива в срезе", sl1, cap(sl1)) // capacity

	sl1 = append(sl1, 101)
	sl1 = append(sl1, 102)
	sl1 = append(sl1, 103)

	fmt.Println("Длина внутреннего массива в срезе", sl1, cap(sl1)) // capacity

	sl1 = append(sl1, 104)

	fmt.Println("Длина внутреннего массива в срезе", sl1, cap(sl1)) // capacity

	fmt.Println("Длина в срезе", sl1, len(sl1))

	// короткая инициализация
	sl2 := []int{
		10,
		20,
		30,
	}
	fmt.Println("Cрез2", sl2)

	// соединение срезов
	var sl3 = append(sl1, sl2...)
	fmt.Println("Cрез3", sl3)

	// набор срезов
	var slm [][]int
	slm = append(slm, sl1)
	fmt.Println("slm", slm)

	// создание среза заданного размера
	sl4 := make([]int, 10)
	fmt.Println(sl4, len(sl4), cap(sl4))
	sl4 = append(sl4, 1)
	fmt.Println(sl4, len(sl4), cap(sl4))

	// создание среза с нужной длиной и размером
	sl5 := make([]int, 10, 15)
	fmt.Println(sl5, len(sl5), cap(sl5))
	sl5 = append(sl5, []int{1,2,3,4,5,6}...)
	fmt.Println(sl5, len(sl5), cap(sl5))

	// внутри среза - ссылка на массив, она копируется если просто присводить
	slice5 := sl4
	slice5[1] = 100500
	fmt.Println(sl4, slice5)

	// правильное копирование массивов среза
	sl7 := make([]int, len(sl5), len(sl5)) // нужен слайс того же размера, что и размер данных для копирования
	copy(sl7, sl5)
	fmt.Println("sl7", sl7)

	// получение части слайса
	fmt.Println("Часть слайса", sl7[1:5], sl7[:2], sl7[10:])

	// создание слайса из массива
	a := [...]int{5,6,7}
	sl8 := a[:]
	a[1] = 8
	fmt.Println("Слайс из массива", sl8)
}
