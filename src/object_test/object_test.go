package object_test

import (
  "testing"
  "fmt"
)

type Employee struct {
  Id string
  Name string
  Age int
}

func TestCreateEmployee(t *testing.T) {
  e := Employee{"0","Bob",20}
  e1 := Employee{Name: "Mike", Age : 30}
  e2 := new(Employee) //返回指针
  e2.Id = "1"
  e2.Name = "Rose"
  e2.Age = 39
  t.Log(e)
  t.Log(e1)
  t.Log(e1.Id)
  t.Log(e2)
  t.Logf("e is %T", e)
  t.Logf("e2 is %T", e2)
}

func (e *Employee) String() string {
  return fmt.Sprintf("Id:%s  Name:%s  Age:%d",e.Id,e.Name,e.Age )
}

// func (e Employee) String() string {
//   return fmt.Sprintf("Id:%s  Name:%s  Age:%d",e.Id,e.Name,e.Age )
// }

func TestStructOperations(t *testing.T){
  e := Employee{"0","Bob",20}
  t.Log(e.String())
}


//txtend  go语言不支持继承 下面是嵌套（类作为属性）的示例
type Pet struct{
}

func (p *Pet) Speak() {
  fmt.Println("...")
}

func (p *Pet) SpeakTo(host string)  {
  p.Speak()
  fmt.Println(host)
}

type Dog struct{
  Pet //匿名嵌套类型
}

// func (d *Dog) Speak() {
//   d.p.Speak()
// }
//
// func (d *Dog) SpeakTo(host string)  {
//   d.p.SpeakTo(host)
// }

func TestExtend(t *testing.T) {
  dog := Dog{}
  dog.Speak()
  dog.SpeakTo("hello")
}
