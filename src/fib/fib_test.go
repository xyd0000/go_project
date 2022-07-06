package fib

import "testing"
import "fmt"

func TestFib(t *testing.T){
  var a,b = 1,1
  fmt.Print(a)
  for i:=0;i<5;i++{
    fmt.Print(" ",b)
    c := a
    a = b
    b = c+a
  }
  fmt.Println()
}
