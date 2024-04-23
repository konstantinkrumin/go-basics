/*
	Современные процессоры (CPU) имеют несколько ядер, и некоторые поддерживают 
	гиперпоточность, поэтому они могут обрабатывать несколько инструкций одновременно. 
	Чтобы полностью утилизировать современные CPU, нужно использовать конкурентное 
	программирование.

	Конкурентные вычисления — это форма вычислений, когда несколько инструкций выполняются, 
	пересекаясь, в течение одного временного периода.

	Например есть 2 инструкции, которые нужно выполнить: А и B. При выполнении инструкций 
	конкурентно ядро процессора вычисляет инструкцию A, при ожидающей B. Когда первая 
	инструкция A заканчивает активное вычисление, она ставится на паузу, и ядро 
	переключается на вычисление инструкции B.

	Рассмотрим простой пример конкурентности в реальной жизни: HTTP запрос к стороннему 
	сервису по сети:
		- вычисление А делает HTTP запрос
		- пока А ждет ответа, начинает выполняться вычисление B
		- когда ответ пришел, А возвращается в работу

	Если в процессоре присутствует более одного ядра, обработка происходит параллельно. 
	Возвращаясь к примеру выше, инструкции A и B будут выполняться независимо в один 
	момент времени.

	Конкурентные вычисления могут происходить в программе, компьютере или сети. В данном 
	курсе рассматривается только уровень программ.
*/

/* ========================= */
/*
	Написать конкурентный код самостоятельно в данном уроке не получится, поэтому 
	давайте подготовимся к следующему уроку реализацией синхронного кода.

	Реализуйте функцию MaxSum(nums1, nums2 []int) []int, которая суммирует числа в каждом 
	слайсе и возвращает слайс с наибольшей суммой. Если сумма одинаковая, то возвращается 
	первый слайс.

	Пример работы:

	MaxSum([]int{1, 2, 3}, []int{10, 20, 50}) // [10 20 50]
	MaxSum([]int{3, 2, 1}, []int{1, 2, 3}) // [3 2 1]
*/
package solution

func MaxSum(nums1, nums2 []int) []int {
	sum1 := 0
	sum2 := 0

	for _, num := range nums1 {
		sum1 += num
	}

	for _, num := range nums2 {
		sum2 += num
	}

	if (sum2 > sum1) {
		return nums2
	}

	return nums1
}

// OR

func MaxSum(nums1, nums2 []int) []int {
	if sum(nums1) >= sum(nums2) {
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