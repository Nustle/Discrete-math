package main

import (
  "bufio"
  "fmt"
  "os"
)

type Stack struct {
  top  int
  data []rune
}

func InitStack(stack *Stack) {
  stack.top = 0
  stack.data = make([]rune, stack.top+1)
}

func Push(stack *Stack, x rune) {
  if stack.top == len(stack.data) {
    stack.data = append(stack.data, x)
  } else {
    stack.data[stack.top] = x
  }
  stack.top++
}

func Pop(stack *Stack) rune {
  stack.top--
  return stack.data[stack.top]
}

func calculate(first, second int, op rune) rune {
  var res int
  if op == '+' {
    res = first + second
  } else if op == '-' {
    res = first - second
  } else {
    res = first * second
  }
  return rune(res) + '0'
}

func calcPolish(s []rune, operandStack, opStack *Stack) int {
  for _, v := range s {
    if v == '(' || v == '+' || v == '*' || v == '-' {
      Push(opStack, v)
    } else if v == ')' {
      op := Pop(opStack)
      for op != '(' {
        second := Pop(operandStack)
        first := Pop(operandStack)
        operand := calculate(int(first-'0'), int(second-'0'), op)
        Push(operandStack, operand)
        op = Pop(opStack)
      }
    } else if v != ' ' {
      Push(operandStack, v)
    }
  }
  ans := Pop(operandStack)
  return int(ans - '0')
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
  ans := calcPolish(s, &operandStack, &opStack)
  fmt.Printf("%d", ans)
}
