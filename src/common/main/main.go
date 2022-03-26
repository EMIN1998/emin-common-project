package main

import (
	"common/trie"
	"fmt"
)

/**
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	var tmpRes, max int
	for _, a := range array {
		if tmpRes < 0 {
			tmpRes = a
		} else {
			tmpRes += a
		}

		if max < tmpRes {
			max = tmpRes
		}

	}

	return max
}

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

func main() {
	//data := make([]int, 5, 10)
	//fmt.Printf("index:%p, cap:%d, len:%d\n", data, cap(data), len(data))
	//for i := 0; i <= 4; i++ {
	//	data = append(data, i)
	//}
	//fmt.Printf("index:%p, cap:%d, len:%d\n", data, cap(data), len(data))
	//for i := 5; i <= 20; i++ {
	//	data = append(data, i)
	//}
	//fmt.Printf("index:%p, cap:%d, len:%d\n", data, cap(data), len(data))
	//arr := data[3:10]
	//fmt.Printf("index:%p, cap:%d, len:%d\n", arr, cap(arr), len(arr))
	//for idx, _ := range arr {
	//	arr[idx] += 10
	//}
	//fmt.Printf("%v", arr)
	//fmt.Printf("%v", data)
	//wg := sync.WaitGroup{}
	//ch := make(chan int)
	//////wg.Add(1)
	//////wg.Add(1)
	//go func(ch chan int) {
	//	//wg.Add(1)
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//		fmt.Printf("in :%d\n", i)
	//	}
	//	//wg.Done()
	//}(ch)
	//
	//for i := 0; i < 10; i++ {
	//	msg := <-ch
	//	fmt.Printf("msg:%d\n", msg)
	//}
	//close(ch)
	//ch <- 1000
	//ms := <-ch
	//
	////go func(ch chan int) {
	////	for i := 0; i < 10; i++ {
	////		msg := <-ch
	////		fmt.Printf("msg:%d\n", msg)
	////	}
	////	//wg.Add(1)
	////	//flag := true
	////	//for flag {
	////	//	select {
	////	//	case msg := <-ch:
	////	//		fmt.Printf("msg :%v\n", msg)
	////	//	default:
	////	//		fmt.Printf("null msg\n")
	////	//		flag = false
	////	//	}
	////	//	//break
	////	//}
	////
	////	fmt.Printf("read done\n")
	////	//wg.Done()
	////}(ch)
	//
	//time.Sleep(time.Second * 5)
	//fmt.Println(ms)
	//fmt.Println("all done ===")
	//cat := testinterface.NewAnimal("Cat")
	//cat.Roar()
	//var a testinterface.Animal
	//a = testinterface.Dog{}
	//a.Roar()
	//fmt.Print()
	//a := testinterface.Dog{Name: "234"}
	////b := testinterface.Dog{Name: "345"}
	//f := testinterface.Dog.Roar
	////fmt.Print()
	//f(a)

	//flag.Parse()
	//fmt.Printf("Sleeping for %v...", *period)
	//time.Sleep(*period)
	//fmt.Println()

	obj := trie.Constructor()
	op := []string{"insert", "search", "search", "startsWith", "insert", "search"}
	//	[[],]
	value := []string{"apple", "apple", "app", "app", "app", "app"}
	resp := make([]*bool, 0)
	for idx, s := range op {
		switch s {
		case "insert":
			obj.Insert(value[idx])
			resp = append(resp, nil)
		case "search":
			r := obj.Search(value[idx])
			resp = append(resp, &r)
		case "startsWith":
			r := obj.StartsWith(value[idx])
			resp = append(resp, &r)
		}
	}

	for _, r := range resp {
		if r == nil {
			fmt.Println("null")
			continue
		}

		p := *r
		fmt.Println(p)
	}
}

//var period = flag.Duration("period", 3*time.Second, "sleep period")
