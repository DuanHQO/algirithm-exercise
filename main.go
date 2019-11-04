package main

import (
	"algorithm/sortfunc"
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


	//list := new(linkedList.List)
	//list.CopyFromArray(origin)
	//list.MergeSort()

	sortfunc.Quick3way(origin, 0, len(origin) - 1)

	fmt.Printf("%v", origin)



	fmt.Println()

}



