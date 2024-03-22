/*
	Условия в Go представлены привычной конструкцией if else. 
	В условии должно быть строго выражение логического типа. 
	Следующий пример вернет ошибку компиляции:
*/
if "hi" { // non-bool "hi" (type string) used as if condition
}

// Корректный пример:

package main

import (
    "fmt"
    "strings"
)

func statusByName(name string) string {
    // функция проверяет, что строка name начинается с подстроки "Mr."
    if strings.HasPrefix(name, "Mr.") {
        return "married man"
    } else if strings.HasPrefix(name, "Mrs.") {
        return "married woman"
    } else {
        return "single person"
    }
}

func main() {
    n := "Mr. Doe"
    fmt.Println(n + " is a " + statusByName(n)) // Mr. Doe is a married man

    n = "Mrs. Berry"
    fmt.Println(n + " is a " + statusByName(n)) // Mrs. Berry is a married woman

    n = "Karl"
    fmt.Println(n + " is a " + statusByName(n)) // Karl is a single person
}

/* ========================= */
/*
	Реализуйте функцию DomainForLocale(domain, locale string) string, которая 
	добавляет язык locale как поддомен к домену domain. Язык может прийти пустым, 
	тогда нужно добавить поддомен en
*/
package solution

import (
	"fmt"
)

func DomainForLocale(domain, locale string) string {
	if (locale == "") {
		return "en." + domain
	} else {
		return locale + "." + domain
	}
}

func main() {
	fmt.Println(DomainForLocale("site.com", "")) // en.site.com
	fmt.Println(DomainForLocale("site.com", "ru")) // ru.site.com
}