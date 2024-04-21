// Возвращаемые ошибки принято проверять при каждом вызове:
import "log"

response, err := DoHTTPCall()
if err != nil {
    log.Println(err)
}
// только после проверки на ошибку можно делать что-то с объектом response

/*
	При этом логика обработки отличается от места и типа ошибки. Ошибки можно оборачивать 
	и прокидывать в функцию выше, логировать или делать любые фоллбек действия.

	Оборачивание ошибок — важная часть написания кода на Go. Это позволяет явно видеть 
	трейс вызова и место возникновения ошибки. Для оборачивания используется функция 
	fmt.Errorf:
*/
package main

import (
    "errors"
    "fmt"
)

// для простоты примера опускаем аргументы запроса и ответа
func DoHTTPCall() error {
    err := SendTCP()
    if err != nil {
        // оборачивается в виде "[название метода]: %w". %w — это плейсхолдер для ошибки
        return fmt.Errorf("send tcp: %w", err)
    }

    return nil
}

var errTCPConnectionIssue = errors.New("TCP connect issue")

func SendTCP() error {
    return errTCPConnectionIssue
}

func main() {
    fmt.Println(DoHTTPCall()) // send tcp: TCP connect issue
}

/*
	В современном Go существуют функции для проверки типов конкретных ошибок. Например, 
	ошибку из примера выше можно проверить с помощью функции errors.Is. В данном случае 
	errTCPConnectionIssue обернута другой ошибкой, но функция errors.Is найдет ее 
	при проверке:
*/
err := DoHTTPCall()
if err != nil {
    if errors.Is(err, errTCPConnectionIssue) {
        // в случае ошибки соединения ждем 1 секунду и пытаемся сделать запрос снова
        time.Sleep(1 * time.Second)
        return DoHTTPCall()
    }

    // обработка неизвестной ошибки
    log.Println("unknown error on HTTP call", err)
}

/*
	errors.Is подходит для проверки статичных ошибок, хранящихся в переменных. Иногда 
	нужно проверить не конкретную ошибку, а целый тип. Для этого используется функция 
	errors.As:
*/
package main

import (
    "errors"
    "log"
    "time"
)

// ошибка подключения к базе данных
type ConnectionErr struct{}

func (e ConnectionErr) Error() string {
    return "connection err"
}

func main() {
    // цикл подключения к БД. Пытаемся 3 раза, если не удалось подсоединиться с первого раза.
    tries := 0
    for {
        if tries > 2 {
            log.Println("Can't connect to DB")
            break
        }

        err := connectDB()
        if err != nil {
            // если ошибка подключения, то ждем 1 секунду и пытаемся снова
            if errors.As(err, &ConnectionErr{}) {
                log.Println("Connection error. Trying to reconnect...")
                time.Sleep(1 * time.Second)
                tries++
                continue
            }

            // в противном случае ошибка критичная, логируем и выходим из цикла
            log.Println("connect DB critical error", err)
        }

        break
    }
}

// для простоты функция всегда возвращает ошибку подключения
func connectDB() error {
    return ConnectionErr{}
}

/*
	Вывод программы спустя 3 секунды:

	Connection error. Trying to reconnect...
	Connection error. Trying to reconnect...
	Connection error. Trying to reconnect...
	Can't connect to DB
*/

/* ========================= */
// Какая-то функция возвращает критичные и некритичные ошибки:
// некритичная ошибка валидации
type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
    return "validation error"
}

// критичные ошибки
var (
    errBadConnection = errors.New("bad connection")
    errBadRequest    = errors.New("bad request")
)

/*
	Реализуйте функцию GetErrorMsg(err error) string, которая возвращает текст ошибки, 
	если она критичная. В случае неизвестной ошибки возвращается строка unknown error:

	GetErrorMsg(errors.New("bad connection")) // "bad connection"
	GetErrorMsg(errors.New("bad request")) // "bad request"
	GetErrorMsg(nonCriticalError{}) // ""
	GetErrorMsg(errors.New("random error")) // "unknown error"
*/
package solution

import (
	"errors"
)

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

func GetErrorMsg(err error) string {
	if errors.As(err, &nonCriticalError{}) {
		return "validation error"
	}

	if errors.Is(err, errBadConnection) {
		return errBadConnection.Error()
	}

	if errors.Is(err, errBadRequest) {
		return errBadRequest.Error()
	}

	return "unknown error"
}