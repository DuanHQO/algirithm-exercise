package main

import (
	"algorithm/linkedList"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//origin1 := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	//origin2 := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	//origin3 := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	origin := rand.Perm(20)
	//origin := []byte {'S', 'H', 'E', 'L', 'L', 'S', 'O', 'R', 'T', 'E', 'X', 'A', 'M', 'P', 'L', 'E',}
	fmt.Println(origin)
	//quickSort.QuickSort(origin, 0, len(origin))
	//sortfunc.InsertionSort(origin)
	//sortfunc.ShellSort(origin)

	list := new(linkedList.List)
	list.CopyFromArray(origin)
	list.MergeSort()

	//fmt.Printf("%v", origin)

	origin1 := []int32{"cz","aa","kz","cd","cc","ab"}
	origin2 := []int32{"kz","ae","aa","ad","ac","al"}
	origin3 := []int32{"bz","bb","aa","bd","bc","kz"}



	fmt.Println()
}



