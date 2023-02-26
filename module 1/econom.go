package main

import (
  "bufio"
  "fmt"
  "os"
)

type Stack struct {
  top  int
  data []string
}

func InitStack(stack *Stack) {
  stack.top = 0
  stack.data = make([]string, stack.top+1)
}

func Push(stack *Stack, x string) {
  if stack.top == len(stack.data) {
    stack.data = append(stack.data, x)
  } else {
    stack.data[stack.top] = x
  }
  stack.top++
}

func Pop(stack *Stack) string {
  stack.top--
  return stack.data[stack.top]
}

func economPolish(s []rune, operandStack, opStack *Stack, history map[string]struct{}) int {
  for _, v := range s {
    if v == '(' || v == '#' || v == '$' || v == '@' {
      Push(opStack, string(v))
    } else if v == ')' {
      op := Pop(opStack)
      for op != "(" {
        second := Pop(operandStack)
        first := Pop(operandStack)
        operand := op + first + second
        Push(operandStack, operand)
        history[operand] = struct{}{}
        op = Pop(opStack)
      }
    } else {
      Push(operandStack, string(v))
    }
  }
  ans := len(history)
  return ans
}

func main() {
  var s []rune
  sc := bufio.NewScanner(os.Stdin)
  sc.Scan()
  s = []rune(sc.Text())
  var operandStack Stack
  InitStack(&operandStack)
  var opStack Stack
  InitStack(&opStack)
  var history = make(map[string]struct{})
  ans := economPolish(s, &operandStack, &opStack, history)
  fmt.Println(ans)
}
