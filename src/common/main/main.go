package main

import (
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
	//slice := make([]int, 0, 10)
	//for i := 0; i < 5; i++ {
	//	slice = append(slice, i)
	//}
	//
	//tmpSlice := make([]int, 0, 10)
	//for i := 0; i < 5; i++ {
	//	tmpSlice = append(slice, i)
	//}
	//fmt.Printf(" tmpSlice after %p, value: %v \n", &tmpSlice, tmpSlice)
	//fmt.Printf("before %p, value : %v\n", &slice, slice)
	//try(slice)
	//
	//fmt.Printf("after %p, value: %v \n", &slice, slice)
	map1 := make(map[int]struct{})
	map2 := make(map[int]struct{})

	for i := 0; i < 10; i++ {
		map1[i] = struct{}{}
	}

	for i := 5; i < 15; i++ {
		map2[i] = struct{}{}
	}

}

func tryDefer(i int) int {
	defer func() {
		i = i + 10
	}()

	return i
}

func try(slice []int) {
	fmt.Printf("try before %p, value: %v \n", &slice, slice)
	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		//atomic.AddInt32()
	}

	for i := 0; i < 5; i++ {
		slice[i] += 100
	}

	fmt.Printf("try %p, value: %v \n", &slice, slice)
}

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
		tmap[k] = "testtest"
		fmt.Println("[", funcName, "]", k, " ==>", v)
	}

	for k, v := range tmap {
		fmt.Println("[", funcName, "]", "after", k, " ==>", v)
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
