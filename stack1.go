//自己练习部分
// func isValid(s string) bool {
//   if ch == "(" || ch == "{" || ch == "[" {
//        stack.push(ch)
//   }
//    if ch == ")" && stack.data[stack.top] =="(" {
//        stack.pop()
//    }
// }

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var stack []rune
	for i, ch := range s {
		if ch == "(" || ch == "{" || ch == "[" {
			stack = append(stack, rune(ch)) //左括号压入栈中
		} else {
			//else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] //最右砍掉，相当于将top弹出
			if (ch == ")" && top != "(") || (ch == "}" && top != "{") || (ch == "]" && top != "[") {
				return false
			}

		}

	}

	return len(stack) == 0
}

func isValid(s string) bool {
	var stack []rune // 声明一个数组作为栈
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, rune(c)) // 压入栈顶元素
		} else {
			if len(stack) == 0 { // 如果第一次压入的是右括号，没有元素进行匹配
				return false
			}

			if len(stack) > 0 {
				topChar := stack[len(stack)-1] // 获取栈顶元素
				stack = stack[:len(stack)-1]   // 弹出栈顶元素
				if c == ')' && topChar != '(' {
					return false
				}
				if c == ']' && topChar != '[' {
					return false
				}
				if c == '}' && topChar != '{' {
					return false
				}
			}
		}
	}
	return len(stack) == 0 // 判断栈是否为空
}

//3
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var stack []rune
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, rune(ch)) //左括号压入栈中
		} else {
			if len(stack) == 0 {
				return false
			}
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1] //最右砍掉，相当于将top弹出

				if (ch == ')' && top != '(') || (ch == '}' && top != '{') || (ch == ']' && top != '[') {
					return false
				}
				// if c == ')' && topChar != '(' {
				// return false
				// }
				// if c == ']' && topChar != '[' {
				// return false
				// }
				// if c == '}' && topChar != '{' {
				// return false
				// }

			}

		}
	}
	return len(stack) == 0
}
