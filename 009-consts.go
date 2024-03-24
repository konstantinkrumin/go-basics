// const [название] [тип данных] = [значение]
const StatusOk int = 200

/*
	На практике тип данных не указывается, и несколько констант принято объявлять в рамках 
	одного блока const:
*/
const (
    StatusOk = 200
    StatusNotFound = 404
)

/*
	Только некоторые типы данных можно присвоить константе: строки, символы, числа, 
	логический тип:
*/
package main

type Person struct {
}

func main() {
    // такие константы допустимы
    const (
        num = 20
        str = "hey"
        isValid = true
    )

    // нельзя объявить структуру как константу
    const p = Person{} // ошибка компиляции: const initializer Person{} is not a constant
}

/*
	Для последовательных числовых констант следует использовать идентификатор iota, 
	который присвоит для списка чисел значения от его текущей позиции:
*/
package main

import "fmt"

const (
    zero = iota
    one
    two
    three
)

const (
	a = iota
	b = 42
	c = iota
	d
)

func main() {
fmt.Println(zero, one, two, three) // 0 1 2 3
fmt.Println(a, b, c, d)            // 0 42 2 3
}

/* ========================= */
/*
	 Реализуйте функцию ErrorMessageToCode(msg string) int, которая возвращает числовой код
	 для заданного значения. Список сообщений и соответствующих кодов:
*/
package solution

func ErrorMessageToCode(msg string) int {
	const (
		OK = 0
		CANCELLED = 1
		UNKNOWN = 2
	)

	switch msg {
		default:
			return UNKNOWN
		case "OK":
			return OK
		case "CANCELLED":
			return CANCELLED
	}
}