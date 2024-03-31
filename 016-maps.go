/*
	Map — тип данных, предназначенный для хранения пар ключ-значение. В других языках 
	эту структуру так же называют: хэш-таблица, словарь, ассоциативный массив. 
	Запись и чтение элементов происходят в основном за O(1):
*/
// создание пустой мапы
var m map[int]string

// сокращенное создание пустой мапы
m := map[int]string{}

// рекомендуемое создание с обозначением размера
m := make(map[int]string, 10)

// создание мапы с элементами
m := map[int]string{1: "hello", 2: "world"}

// добавление элемента
m[3] = "!" // map[1:hello, 2:world, 3:!]

// чтение элемента
word := m[1] // "hello"

/*
	При чтении элемента по несуществующему ключу возвращается нулевое значение данного типа.
	Это приводит к ошибкам логики, когда используется bool как значение. Для решения 
	данной проблемы при чтении используется вторая переменная, в которую записывается 
	наличие элемента в мапе:
*/
elements := map[int64]bool{1: true, 2: true}

element, elementExists := elements[1] // true, true
element, elementExists := elements[2] // true, true

element, elementExists := elements[225] // false, false

// Для проверки существования ключа можно использовать мапу с пустыми структурами
// пустая структура struct{} — это тип данных, который занимает 0 байт
// используется, когда нужно проверять в мапе только наличие ключа
cache := make(map[string]struct{})

// проверяем есть ли ключ `key` в мапе
_, ok = cache["key"]
fmt.Println(ok)  // false

// добавим ключ и проверим вновь
cache["key"] = struct{}{}
_, ok = cache["key"]
fmt.Println(ok)  // true

// Элементы удаляются с помощью встроенной функции delete(m map[Type]Type1, key Type):
engToRus := map[string]string{"hello":"привет", "world":"мир"}

delete(engToRus, "world")

fmt.Println(engToRus) // map[hello:привет]

// Мапы в Go всегда передаются по ссылке:
package main

import (
    "fmt"
)

func main() {
    m := map[int]string{1: "hello", 2: "world"}

    modifyMap(m)

    fmt.Println(m) // вывод: map[1:changed 2:world 200:added]
}

func modifyMap(m map[int]string) {
    m[200] = "added"

    m[1] = "changed"
}

/* ========================= */
/*
	Реализуйте функцию UniqueUserIDs(userIDs []int64) []int64, которая возвращает слайс, 
	состоящий из уникальных идентификаторов userIDs. Порядок слайса должен сохраниться.
*/
func UniqueUserIDs(userIDs []int64) []int64 {
	// пустая структура struct{} — это тип данных, который занимает 0 байт
	// используется, когда нужно проверять в мапе только наличие ключа
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, uid := range userIDs {
		_, is_exists := processed[uid]
		if is_exists {
			continue
		}

		uniqUserIDs = append(uniqUserIDs, uid)
		processed[uid] = struct{}{}
	}

	return uniqUserIDs
}