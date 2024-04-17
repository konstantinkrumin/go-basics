/*
	В Go нет классов, но существуют структуры с методами. Метод — это функция с 
	дополнительным аргументом, который указывается в скобках между func и названием функции:
*/
package main

import (
    "fmt"
)

type Dog struct{}

/* 
	сначала объявляется дополнительный аргумент "(d Dog)", 
	а следом идет обычное описание функции
*/
func (d Dog) Bark() {
    fmt.Println("woof!")
}

func main() {
    d := Dog{}
    d.Bark() // woof!
}

/*
	В примере выше структура Dog передается по значению, то есть копируется. 
	Если изменятся любые свойства внутри метода Bark, они останутся неизменными в 
	исходной структуре:
*/
package main

import (
    "fmt"
)

type Dog struct {
    IsBarked bool
}

func (d Dog) Bark() {
    fmt.Println("woof!")
    d.IsBarked = true
}

func main() {
    d := Dog{}
    d.Bark() // woof!

    fmt.Println(d.IsBarked) // false
}

/*
	Если есть необходимость в изменении состояния, структура должна 
	передаваться указателем:
*/
package main

import (
    "fmt"
)

type Dog struct {
    IsBarked bool
}

func (d *Dog) Bark() {
    fmt.Println("woof!")
    d.IsBarked = true
}

func main() {
    d := &Dog{}
    d.Bark() // woof!

    fmt.Println(d.IsBarked) // true
}

/* ========================= */
/*
	Реализуйте методы структуры Counter, представляющую собой счётчик, хранящий 
	неотрицательное целочисленное значение и позволяющий это значение изменять:

	- метод Inc(delta int) должен увеличивать текущее значение на delta единиц 
	(на 1 по умолчанию),
	- метод Dec(delta int) должен уменьшать текущее значение на delta единиц.

	c := Counter{}
	c.Inc(0)
	c.Inc(0)
	c.Inc(40)
	c.Value // 42

	c.Dec(0)
	c.Dec(30)
	c.Value // 11

	c.Dec(100)
	c.Value // 0
*/
package solution

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (c *Counter) Inc(delta int) {
	if delta == 0 {
		delta = 1
	}

	c.Value = Max(c.Value+delta, 0)
}


func (c *Counter) Dec(delta int) {
	if delta == 0 {
		delta = 1
	}

	c.Inc(-delta)
}