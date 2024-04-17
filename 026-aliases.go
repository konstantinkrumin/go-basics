/*
	В Go можно объявить алиас на существующий тип данных для выразительности и абстракции. 
	Например, тип byte из модуля "строки" — это алиас uint8. Алиас объявляется через 
	ключевое слово type:
*/
type NumCount int

func main() {
    nc := NumCount(len([]int{1, 2, 3}))

    fmt.Println(nc) // 3
}

// Алиас можно конвертировать в оригинальный тип и обратно:
type errorCode string

func main() {
    ec := errorCode("internal")

    fmt.Println(ec) // internal

    fmt.Println(string(ec)) // internal
}

/*
	Также у алиасов могут быть методы. Объявление метода происходит так же, как и со 
	структурами:
*/
type counter int


// передается указатель, чтобы можно было изменить состояние счетчика "c"
func (c *counter) inc() {
    *c++
}

func main() {
    c := counter(0)
    (&c).inc() // передается указатель на счетчик &c, так как функция "inc()" работает с указателями
    (&c).inc()

    fmt.Println(c) // 2
}

/* ========================= */
/*
	Представим, что есть структура Person, содержащая возраст человека:

	type Person struct {
		Age uint8
	}

	Реализуйте тип PersonList (слайс структур Person), с методом 
	(pl PersonList) GetAgePopularity() map[uint8]int, 
	который возвращает мапу, где ключ — возраст, а значение — кол-во таких возрастов:

	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}

	pl.GetAgePopularity() // map[18:2 44:1]
*/
package solution

// Person is a struct that keeps info about person's age
type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	ageCount := make(map[uint8]int)

	for _, p := range pl {
		age := p.Age
		count, exists := ageCount[age]
		if (exists) {
			ageCount[age] = count + 1
		} else {
			ageCount[age] = 1
		}
	}

	return ageCount
}

// OR

func (pl PersonList) GetAgePopularity() map[uint8]int {
	popularity := make(map[uint8]int)
	for _, p := range pl {
		popularity[p.Age]++
	}

	return popularity
}