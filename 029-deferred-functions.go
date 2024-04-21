/*
	В Go есть полезная конструкция defer, которая позволяет выполнять функции в фазе выхода 
	из текущей функции. Например:
*/
package main

import (
    "fmt"
)

func main() {
    // функция выполнится в самом конце при выходе из main
    defer fmt.Println("finish")

    fmt.Println("start")
}
/*
	Вывод:

		start
		finish
*/

/*
	Такие функции называются отложенными. Каждая такая функция добавляется в стек 
	отложенных функций и будет выполнена в порядке LIFO (Last In First Out):
*/
package main

import (
    "fmt"
)

func main() {
    defer fmt.Println("3rd")
    defer fmt.Println("2nd")

    fmt.Println("1st")
}
/*
	Вывод:

		1st
		2nd
		3rd
*/

/*
	Использование отложенных функций достаточно распространено. Например:
	- закрытие дескриптора файла после работы
	- возвращение соединения с базой данных в общий пул после чтения всех строк
	- закрытие TCP соединения после полного прочтения тела ответа
*/

/* ========================= */
/*
	Реализуйте функцию ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error), 
	которая выполняет джобу MergeDictsJob и возвращает ее. Алгоритм обработки джобы следующий:
		- перебрать по порядку все словари job.Dicts и записать каждое ключ-значение в 
		  результирующую мапу job.Merged
		- если в структуре job.Dicts меньше 2-х словарей, возвращается ошибка 
		  errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
		- если в структуре job.Dicts встречается словарь в виде нулевого значения nil, 
		  то возвращается ошибка errNilDict = errors.New("nil dictionary")
		- независимо от успешного выполнения или ошибки в возвращаемой структуре 
		  MergeDictsJob поле IsFinished должно быть заполнено как true

	Пример работы:
		ExecuteMergeDictsJob(&MergeDictsJob{}) // &MergeDictsJob{IsFinished: true}, "at least 2 dictionaries are required"
		ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"},nil}}) // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b"},nil}}, "nil dictionary"
		ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"},{"b": "c"}}}) // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b", "b": "c"}}}, nil
*/
package solution

import "errors"

// MergeDictsJob is a job to merge dictionaries into a single dictionary.
type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func (j *MergeDictsJob) setFinished() {
	j.IsFinished = true
}

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer job.setFinished()

	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}

	job.Merged = make(map[string]string)

	for _, dict := range job.Dicts {
		if dict == nil {
			return job, errNilDict
		}

		for k, v := range dict {
			job.Merged[k] = v
		}
	}

	return job, nil
}