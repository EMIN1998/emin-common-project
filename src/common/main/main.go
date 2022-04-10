package main

import (
	"common/linklist"
	"fmt"
	"sync"
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
	//funcName := "mainfunc"
	//slice := make([]int, 0)
	//for i := 0; i < 10; i++ {
	//	slice = append(slice, i)
	//}
	//
	//tmap := make(map[int]string)
	//for i := 0; i < 5; i++ {
	//	tmap[i] = fmt.Sprintf("mmm + %d", i)
	//}
	//
	//fmt.Printf("[%s] before %p, %v\n", funcName, slice, slice)
	//fmt.Printf("[%s] before %p\n", funcName, tmap)
	//
	//testfunc(slice, tmap)
	//fmt.Printf("[%s] %p, %v\n", funcName, slice, slice)
	//fmt.Printf("[%s] %p\n", funcName, tmap)
	//for k, v := range tmap {
	//	fmt.Println(k, " ==>", v)
	//	tmap[k] = "test"
	//}

	//fmt.Println(testinterface(nil))
	//nestedDo()
	linklist.Testdectfunc()
}

func testinterface(p interface{}) bool {
	if p == nil {
		fmt.Println("==========")
		return true
	}

	return false
}

//func nestedDo() {
//	once := &sync.Once{}
//	once.Do(func() {
//		once.Do(func() {
//			fmt.Println("test nestedDo")
//		})
//	})
//}

func panicDo() {
	once := &sync.Once{}
	defer func() {
		if err := recover(); err != nil {
			once.Do(func() {
				fmt.Println("run in recover")
			})
		}
	}()
	once.Do(func() {
		panic("panic i=0")
	})

}

func testfunc(slice []int, tmap map[int]string) {
	funcName := "testfunc"

	for i := 11; i < 15; i++ {
		slice = append(slice, i)
	}

	for k, v := range tmap {
		fmt.Println(k, " ==>", v)
		tmap[k] = "test"
	}

	for i := 0; i < 5; i++ {
		slice[i] = slice[i] * 100
	}

	fmt.Printf("[%s] %p, %v\n", funcName, slice, slice)
	fmt.Printf("[%s] %p\n", funcName, tmap)

	return
}

func nestedDo() {
	once1 := &sync.Once{}
	once2 := &sync.Once{}
	once1.Do(func() {
		once2.Do(func() {
			fmt.Println("test nestedDo")
		})
	})
}
