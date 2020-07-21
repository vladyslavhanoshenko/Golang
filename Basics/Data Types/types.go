package main

type UserId int

func main() {
	idx := 1

	var uid UserId = 42

	// даже если базовый тип одинаковый, разные типы несовместимы
	// cannot use uid (type UserId) as type int64 in assignment
	// myId := idx

	myId := UserId(idx)

	println(uid, myId)
}
