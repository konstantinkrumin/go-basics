/*
	В Go нет классов и привычной реализации ООП. Вместо классов в языке используются 
	структуры — наборы полей, имеющих название и тип данных. Объявление структуры имеет 
	следующий вид:
*/
type Person struct {
    // [название поля] [тип данных]
    Name string
    Age int
}

func main() {
    p := Person{Name: "John", Age: 25}

    p.Name // "John"
    p.Age // 25
}

/*
	Структуру можно инициализировать, не передавая значения. В этом случае каждое поле 
	примет свое «нулевое» значение:
*/
func main() {
    p := Person{}

    p.Name // ""
    p.Age // 0
}

/*
	Регистр первой буквы в названии структуры и полей означает публичность точно так же, 
	как в переменных и функциях. Если первая буква заглавная, то структуру можно 
	инициализировать во внешних пакетах. В противном случае она доступна только 
	в рамках текущего пакета:
*/
type Person struct { // структура публична
    Name string // поле публично

    wallet wallet // поле приватно: можно обращаться только внутри текущего пакета
}

type wallet struct { // структура приватна: можно инициализировать только внутри текущего пакета
    id string
    moneyAmount float64
}

/*
	У любого поля структуры можно указать теги. Они используются для метаинформации о поле 
	для сериализации, валидации, маппинга данных из БД и тд. Тег указывается после типа 
	данных через бектики:
*/
type User struct {
    ID int64 `json:"id" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    FirstName string `json:"first_name" validate:"required"`
}

/*
	Тег json используется для названий полей при сериализации/десериализации структуры 
	в json и обратно:
*/
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID        int64  `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
}

func main() {
    u := User{}
    u.ID = 22
    u.Email = "test@test.com"
    u.FirstName = "John"

    bs, _ := json.Marshal(u)

    fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}
}

/*
	Тег validate используется Go-валидатором. В следующем примере присутствует вызов 
	функции у структуры v.Struct(u). Функции структур — методы — мы разберем подробно 
	в соответствующем уроке, а пока просто посмотрите, как происходит вызов:
*/
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID        int64  `validate:"required"`
    Email     string `validate:"required,email"`
    FirstName string `validate:"required"`
}

func main() {
    // создали пустую структуру, чтобы проверить валидацию
    u := User{}

    // создаем валидатор
    v := validator.New()

    // метод Struct валидирует переданную структуру и возвращает ошибку `error`, если какое-то поле некорректно
    fmt.Println(v.Struct(u))
}

/*
	Вывод программы:
	Key: 'User.ID' Error:Field validation for 'ID' failed on the 'required' tag
	Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag
	Key: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag
*/

/* ========================= */
/*
	Реализуйте функцию Validate(req UserCreateRequest) string, которая валидирует запрос и 
	возвращает строку с ошибкой "invalid request", если модель невалидна. Если запрос 
	корректный, то функция возвращает пустую строку. Правила валидации и структура 
	модели описаны ниже. Не используйте готовые библиотеки и опишите правила 
	самостоятельно.
*/
package solution

import (
	"strings"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

func Validate(req UserCreateRequest) string {
	error := "invalid request"

	if (req.FirstName == "" || strings.Contains(req.FirstName, " ")) {
		return error
	}

	if (req.Age <= 0 || req.Age > 150) {
		return error
	}

	return ""
}
