package gro

import (
  "testing"
  "fmt"
  "time"
)

func AsyncService1() chan string{
  var retCh = make(chan string, 1)
  go func(){
    ret := service()
    fmt.Println("returnd result")
    retCh <- ret
    fmt.Println("service exited")
  }()
  return retCh
}

func TestSelect(t *testing.T) {
  select {
  case ret := <- AsyncService1():
    t.Log(ret)
  case <- time.After(time.Millisecond * 100):
    t.Error("time out")
  }
}
