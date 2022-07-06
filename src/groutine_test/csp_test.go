package gro
import (
  "testing"
  "time"
  "fmt"
)

func service() string {
  time.Sleep(time.Millisecond * 50)
  return "Done"
}

func otherTask() {
  fmt.Println("working on something else")
  time.Sleep(time.Millisecond * 100)
  fmt.Println("Task is done")
}

func TestService(t *testing.T) {
  t.Log(service())
  otherTask()
}


func AsyncService() chan string{
  var retCh = make(chan string, 1)
  go func(){
    ret := service()
    fmt.Println("returnd result")
    retCh <- ret
    fmt.Println("service exited")
  }()
  return retCh
}

func TestAsynService(t *testing.T){
  retCh := AsyncService()
  otherTask()
  fmt.Println(<-retCh)
}
