package gro

import (
  "fmt"
  "testing"
  "time"
  "sync"
)

func TestGroutine(t *testing.T) {
  for i := 0;i<10; i++ {
    go func (i int){ //起10个协程，执行是无需的
      fmt.Println(i) //每个协程都是for 语句中i的拷贝
    }(i)
  }
}


func TestCounter(t *testing.T) {
  var mut sync.Mutex
  counter := 0
  for i := 0;i<5000; i++{
    go func() {
      defer func(){  //使用defer函数释放锁
        mut.Unlock()
      }()
      mut.Lock()
      counter++  //各协程共享同一个counter，需要通过锁来保证不能同时访问
    }()
  }
  time.Sleep(1*time.Second)
  t.Logf("counter = %d", counter)
}


func TestCounterWaitGroup(t *testing.T) {
  var mut sync.Mutex
  var wg sync.WaitGroup
  counter := 0
  for i := 0;i<5000; i++{
    wg.Add(1)
    go func() {
      defer func(){
        mut.Unlock()
      }()
      mut.Lock()
      counter++
      wg.Done()
    }()
  }
  wg.Wait() //阻塞于协程执行完毕
  t.Logf("counter = %d", counter)
}
