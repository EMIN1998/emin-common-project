package main

import "fmt"

func main() {
	var req = []int{3, 2, 1, 5, 6, 7}
	res := heapSort(req)
	fmt.Printf("\n %v", res)
}
