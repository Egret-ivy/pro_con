func evalRPN(tokens []string) int {

    stack := make([]int,len(tokens))
    top:=0
    for _,ch := range tokens {
        switch ch {
            case "+":
                stack[top-2]=stack[top-2]+stack[top-1]
                top--
            case "-":
                stack[top-2]=stack[top-2]-stack[top-1]
                top--
            case "*":
                stack[top-2]=stack[top-2]*stack[top-1]
                top--
            case "/":
                stack[top-2]=stack[top-2]/stack[top-1]
                top--
            default:
                num,_:=strconv.Atoi(ch)
                stack[top]=num
                top++
        }
    }
    return stack[top-1]
}