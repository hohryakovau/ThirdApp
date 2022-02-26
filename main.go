package main

import (
	"fmt"
	"sort"
)

func main() {
	var posl []int = []int{4, 1, 3, 2}
	//var posl []int = []int{4, 1, 3}  //для тестирования

	fmt.Println(Solution(posl))
	//Функция возвращает 1 если массив А это последовательность и 0 если нет.
}

func Solution(A []int) int {

	sort.Ints(A)
	for i := 0; i < len(A)-1; i++ {
		if A[i] != (A[i+1] - 1) {
			return 0
		}
	}
	return 1

}

//последовательность
//A[0] = 4
//A[1] = 1
//A[2] = 3
//A[3] = 2
//
//
//не последовательность
//A[0] = 4
//A[1] = 1
//A[2] = 3
