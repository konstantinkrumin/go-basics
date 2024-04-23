/*
	Вот и подошло время познакомиться с самой сильной стороной языка Go — горутинами. 
	Горутины — это легковесные потоки, которые реализуют конкурентное программирование в Go. 
	Их называют легковесными потоками, потому что они управляются рантаймом языка, 
	а не операционной системой. Стоимость переключения контекста и расход памяти 
	намного ниже, чем у потоков ОС. Следовательно, для Go — не проблема поддерживать 
	одновременно десятки тысяч горутин.

	Запустить функцию в горутине — супер легко. Для этого достаточно написать слово go 
	перед вызовом функции:
*/
package main

import (
    "fmt"
    "time"
)

func main() {
    // выведет сообщение в горутине
    go fmt.Println("Hello concurrent world")

    // если не подождать, то программа закончится, не успев, вывести сообщение
    time.Sleep(100 * time.Millisecond)
}

/*
	При написании конкурентного кода возникают новые моменты, которые нужно учитывать: 
	состояние гонки, блокировки, коммуникация между горутинами. Пример программы, 
	которая работает не так, как ожидается:
*/
package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 5; i++ {
        go func() {
            fmt.Println(i)
        }()
    }

    time.Sleep(100 * time.Millisecond)
}
/*
	Сперва может показаться, что должны вывестись числа от 0 до 4, но на самом деле вывод 
	будет следующим:

	5
	5
	5
	5
	5
*/

/*
	Все потому что i передается в общем скоупе, следовательно, когда горутины будут 
	выполняться, цикл уже закончится и i будет равно 5. В данном случае нужно передать 
	копию i:
*/
package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 5; i++ {
        go func(i int) {
            fmt.Println(i)
        }(i)
    }

    time.Sleep(100 * time.Millisecond)
}
/*
	Вывод:

	0
	4
	3
	1
	2
*/

/*
	Также можно заметить, что числа вывелись не в порядке вызова. Горутины выполняются 
	независимо и не гарантируют порядка. При необходимости последовательность в выполнении 
	придется реализовывать самостоятельно.
*/

/* ========================= */
/*
	Реализуйте функцию MaxSum(nums1, nums2 []int) []int из прошлого задания, используя 
	горутины для расчета каждой суммы слайса.

	Не забудьте использовать функцию time.Sleep(100 * time.Millisecond), чтобы сумма успела 
	посчитаться. В настоящих приложениях используются специальные инструменты, чтобы 
	дожидаться исполнения асинхронного кода, но для простоты здесь будет использоваться 
	обычный сон.
*/
package solution

import (
	"time"
)

func MaxSum(nums1, nums2 []int) []int {
	sum1 := 0
	sum2 := 0

	go func() {
		sum1 = sum(nums1)
	}()

	go func() {
		sum2 = sum(nums2)
	}()


	time.Sleep(100 * time.Millisecond)

	if sum1 >= sum2 {
		return nums1
	}

	return nums2
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	return s
}

// OR

// MaxSum gets sum of each nums slice and returns a slice with the max sum.
// If the slices are the same then the first one will be returned.
func MaxSum(nums1, nums2 []int) []int {
	s1, s2 := 0, 0

	go sumParallel(nums1, &s1)
	go sumParallel(nums2, &s2)

	time.Sleep(100 * time.Millisecond)

	if s1 >= s2 {
		return nums1
	}

	return nums2
}

func sumParallel(nums []int, res *int) {
	*res = sum(nums)
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	return s
}