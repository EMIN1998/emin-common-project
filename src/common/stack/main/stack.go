package main

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	tmpStack := *s
	if len(tmpStack) == 0 {
		return nil
	}

	*s = tmpStack[:tmpStack.Len()-1]
	// res := tmpStack[tmpStack.Len()-1]
	return tmpStack[tmpStack.Len()-1]
}

func (s *Stack) Top() interface{} {
	if len(*s) == 0 {
		return nil
	}

	tmpStack := *s
	return tmpStack[len(tmpStack)-1]
}

// func teststack() {
// 	s := make(Stack, 0)
// }
