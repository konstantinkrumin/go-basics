/*
	В предыдущих уроках мы использовали пакет fmt для вывода переменных или результатов 
	функций:
*/
s := "hello world"

// печатает вывод на следующей строке
fmt.Println(s) // "hello world"

/*
	Пакет fmt так же используется для форматирования строк. Плейсхолдеры разных типов 
	данных в основном не отличаются от других языков:
*/
name := "Andy"

// подставляем строку
fmt.Sprintf("hello %s", name) // "hello Andy"

// число
fmt.Sprintf("there are %d kittens", 10) // "there are 10 kittens"

// логический тип
fmt.Sprintf("your story is %t", true) // "your story is true"

/*
	Так же существуют специализированные плейсхолдеры, которые преобразуют 
	сложные структуры:
*/
package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Andy", Age: 18}

    // вывод значений структуры
    fmt.Println("simple struct:", p)

    // вывод названий полей и их значений
    fmt.Printf("detailed struct: %+v\n", p)

    // вывод названий полей и их значений в виде инициализации
    fmt.Printf("Golang struct: %#v\n", p)
}
/*
	Вывод:
		simple struct: {Andy 18}
		detailed struct: {Name:Andy Age:18}
		Golang struct: main.Person{Name:"Andy", Age:18}
*/

/* ========================= */
/*
	Реализуйте функцию generateSelfStory(name string, age int, money float64) string, 
	которая генерирует предложение, подставляя переданные данные в возвращаемую строку. 
	Например:

	generateSelfStory("Vlad", 25, 10.00000025) // "Hello! My name is Vlad. I'm 25 y.o. 
	And I also have $10.00 in my wallet right now."

	Шаблон возвращаемой строки: Hello! My name is [name]. I'm [age] y.o. And I also have 
	$[money with precision 2] in my wallet right now.
*/
package solution

import "fmt"

func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}