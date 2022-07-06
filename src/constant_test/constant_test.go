package constant_test

import "testing"

const (
  Monday = 1 + iota
  Tuesday
)

const (
  Readable = 1 << iota
  Writable
  Executable
)

func TestConstantTry(t *testing.T){
  t.Log(Monday,Tuesday)
  a := 7 //0111
  t.Log(a&Readable==Readable,a&Writable==Writable,a&Executable == Executable)
}
