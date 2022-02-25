package main

import "fmt"

/**
 *
 * @param A int整型一维数组
 * @param B int整型一维数组
 * @return void
 */
func merge(A []int, m int, B []int, n int) {
	// write code here
	var i, j, p int
	i = m - 1
	j = n - 1
	p = m + n - 1

	for {
		if p < 0 || j < 0 {
			break
		}

		if i < 0 {
			A = append(A, B...)
			break
		}
		if A[i] < B[j] {
			A[p] = B[j]
			p--
			j--
		} else {
			A[p] = A[i]
			p--
			i--
		}

		if i == p {
			break
		}

		if i < 0 && j >= 0 {
			for {
				A[p] = B[j]
				j--
				p--
				if p < 0 {
					break
				}
			}
		}

		if j < 0 && i >= 0 {
			for {
				A[p] = A[i]
				i--
				p--
				if p < 0 {
					break
				}
			}
		}
	}

	fmt.Print(A)
	return
}

func main() {
	a := []int{}
	b := []int{1}
	merge(a, 0, b, 1)
}
