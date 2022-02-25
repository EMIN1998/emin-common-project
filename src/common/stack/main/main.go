package main

import (
	"fmt"
	"strconv"
)

func solve(input string) int64 {
	req := splitByOpt(input)
	numStack := make(Stack, 0)
	optStack := make(Stack, 0)
	tmpOpt := ""
	for _, s := range req {
		if string(s) == "(" {
			if tmpOpt != "" {
				optStack.Push(tmpOpt)
				tmpOpt = ""
			}

			optStack.Push(s)
			continue
		}

		if tmpOpt != "" {
			switch tmpOpt {
			case "-":
				tmps := fmt.Sprintf("%s%s", tmpOpt, s)

				tmpNum, _ := strconv.ParseInt(tmps, 10, 64)
				numStack.Push(tmpNum)
				optStack.Push("+")
			default:
				if s == "(" {
					optStack.Push(tmpOpt)
				} else {
					num, _ := strconv.ParseInt(s, 10, 32)
					num1 := numStack.Pop().(int64)
					resp := calculate(num, num1, tmpOpt)
					numStack.Push(resp)
				}

			}

			tmpOpt = ""
			continue
			//numStack.Push()
		}

		if s == ")" {
			a := numStack.Pop().(int64)
			b := numStack.Pop().(int64)
			o := optStack.Pop().(string)

			lastResp := calculate(b, a, o)
			for {
				opt := optStack.Pop().(string)
				if opt == "(" {
					if optStack.Len() != 0 {
						op := optStack.Pop().(string)
						if isSeniorOpt(op) {
							switch tmpOpt {
							case "-":
								tmps := fmt.Sprintf("%s%s", tmpOpt, s)

								tmpNum, _ := strconv.ParseInt(tmps, 10, 64)
								numStack.Push(tmpNum)
								optStack.Push("+")
							default:
								{

									num := numStack.Pop().(int64)
									lastResp = calculate(num, lastResp, op)
									//numStack.Push(resp)
								}

							}
						} else {
							optStack.Push(op)
						}
					}
					numStack.Push(lastResp)
					break
				}

				lastNum := numStack.Pop().(int64)
				lastResp = calculate(lastNum, lastResp, opt)
			}

			continue
		}

		num, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			if isSeniorOpt(s) {
				tmpOpt = s
				continue
			}
			optStack.Push(s)
		} else {

			numStack.Push(num)
		}

	}

	//result := 0
	for {
		if len(numStack) == 1 {
			return numStack.Pop().(int64)
		}

		num1 := numStack.Pop().(int64)
		num2 := numStack.Pop().(int64)
		result := calculate(num2, num1, optStack.Pop().(string))
		numStack.Push(result)
	}

}

func splitByOpt(s string) []string {
	resp := make([]string, 0)
	tmpNum := ""
	for _, v := range s {
		switch string(v) {
		case "+", "-", "*", "/", "(", ")":
			if tmpNum != "" {
				resp = append(resp, tmpNum)
				tmpNum = ""
			}

			resp = append(resp, string(v))
		default:
			tmpNum += string(v)
		}
	}

	if tmpNum != "" {
		resp = append(resp, tmpNum)
	}
	return resp
}

func isSeniorOpt(s string) bool {
	if s == "*" || s == "-" {
		return true
	} else {
		return false
	}
}

func calculate(a, b int64, op string) int64 {
	switch op {
	case "*":
		return a * b
	case "+":
		return a + b
	case "-":
		return a - b
	default:
		return 0
	}
}

func main() {
	//_ =
	fmt.Println(solve("(3+4)*(5+(2-3))"))
}
