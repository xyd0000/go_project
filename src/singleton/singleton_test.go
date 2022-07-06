package singleton

import (
  "testing"
  "sync"
  "fmt"
  "unsafe"
)

type Singleton struct{

}

var singletonInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton{
  once.Do(func(){
    fmt.Println("Create Obj")
    singletonInstance = new(Singleton)
  })
  return singletonInstance
}

func TestGetSingletonObj(t *testing.T)  {
  var wg sync.WaitGroup
  for i := 1;i<10;i++{
    wg.Add(1)
    go func ()  {
      obj := GetSingletonObj()
      fmt.Printf("%p\n" , obj)
      fmt.Printf("%p\n" , unsafe.Pointer(obj))
      wg.Done()
    }()
  }

}
