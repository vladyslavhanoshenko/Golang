package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Сумма равна ", add(10, 5))

	//короткое присваивание
	//хотя бы одна переменная должна быть новой
	var isTrue, isFalse bool

	isTrue = false
	isFalse = true

	fmt.Println(isTrue)
	fmt.Println(isFalse)

	var summ int

	for i := 0; i < 10; i++ {
		summ += i
	}

	fmt.Println(summ)

	//короткое объявление переменнной
	//только для новых переменных
	//принято использовать camelCase
	myVariable := 30

	fmt.Println("My new variable is ", myVariable)

	var balance [10]int

	balance[0] = 1
	balance[1] = 2
	balance[2] = 3
	balance[3] = 4

	fmt.Println(balance)
	fmt.Println(balance[0:2])

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	V := Vertex{1, 2}
	V.X = 4
	fmt.Println("_______________")
	fmt.Println(V.X)

}

type Vertex struct {
	X int
	Y int
}
