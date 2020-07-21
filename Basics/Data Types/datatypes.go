package main

import "fmt"

func main() {

	// платформозависимый тип, 32/64
	var i int = 10

	// int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// платформозависимый тип, 32/64
	var unsignedInt uint = 100500

	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(i, bigInt, unsignedInt, unsignedBigInt)

	// float32, float64
	var pi float32 = 3.141
	var e = 2.718
	goldenRation := 1.618

	fmt.Println(pi, e, goldenRation)

	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 + 7.12i

	fmt.Println(c, c2)

	//пустая строка
	var str string
	fmt.Println(str)

	//строки поддерживают UTF-8 из коробки
	//строки неизменяемые

	//rune (uint32) для UTF-8 символов
	var someChinese rune = '漢'
	fmt.Println(someChinese)

	newString := "My custom new string"
	fmt.Println(len(newString))
	fmt.Println(newString[:10])

	var newStringNew string
	fmt.Println(newStringNew)

	//конвертация в слайс байт и обратно
	byteString := []byte(newString)
	newStringNew = string(byteString)

	fmt.Println(byteString, newString)
}

//множественное объявление констант
const (
	hello = "Hello" //нетипизированная константа
	e     = 2.718
)

//в go нет автоматического приведения типов
