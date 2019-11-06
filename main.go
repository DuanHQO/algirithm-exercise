package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//origin1 := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	//origin2 := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	//origin3 := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	//origin := rand.Perm(20)
	origin := []int{3, 2, 4}
	//origin := []byte {'S', 'H', 'E', 'L', 'L', 'S', 'O', 'R', 'T', 'E', 'X', 'A', 'M', 'P', 'L', 'E',}
	fmt.Println(origin)

	fmt.Println(twoSum(origin, 6))

	//fmt.Println(origin)

	//quickSort.QuickSort(origin, 0, len(origin))
	//sortfunc.InsertionSort(origin)
	//sortfunc.ShellSort(origin)

	//list := new(linkedList.List)
	//list.CopyFromArray(origin)
	//list.MergeSort()

	//sortfunc.Quick3way(origin, 0, len(origin) - 1)

	//fmt.Printf("%v", origin)

	//Ch_2_5_Applications.TestTrade()

	fmt.Println()

}

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for key, value := range nums {
		hashmap[value] = key
	}
	for i := 0; i < len(nums); i++ {
		offset := target - nums[i]
		value, exists := hashmap[offset]
		if exists && value != i {
			return []int{i, value}
		}
	}
	return nil
}
