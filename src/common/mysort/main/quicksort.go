package main

func quicksort(req *[]int) {
	if len(*req) == 1 || len(*req) == 0 {
		return
	}
	// index := low
	low := 0
	high := len(*req) - 1
	tmp := (*req)[low]
	for {
		if low == high {
			break
		}
		if tmp > (*req)[high] {
			(*req)[low] = (*req)[high]
		} else {
			high--
			continue
		}

		if tmp < (*req)[low] {
			(*req)[high] = (*req)[low]
		} else {
			low++
			continue
		}
	}

	(*req)[low] = tmp
	pre := (*req)[0:low]
	quicksort(&pre)

	sub := (*req)[low+1 : len(*req)]
	quicksort(&sub)
}
