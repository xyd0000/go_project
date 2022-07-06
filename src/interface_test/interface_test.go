package interface_test
import "testing"
import "fmt"


type Programmer interface {
  WriteHelloWorld( ) string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld( ) string {
  return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld( ) string {
  return "System.out.Println(\"Hello World\")"
}

func writeFirstProgram(p Programmer) {
  fmt.Printf("%T %v\n",p ,p.WriteHelloWorld())
}
func TestClient(t *testing.T) {
  var p Programmer
  p = new(GoProgrammer)
  t.Log(p.WriteHelloWorld())
  //多态
  gop := new(GoProgrammer)
  javap := new(JavaProgrammer)
  writeFirstProgram(gop)
  writeFirstProgram(javap)
}


func DoSomething(p interface{}){
  // if i,ok := p.(int); ok{
  //   fmt.Println("Integer",i)
  //   return
  // }
  // if s,ok := p.(string); ok{
  //   fmt.Println("string",s)
  //   return
  // }
  // fmt.Println("Unknow Type")
  switch v:=p.(type) {
  case int:
    fmt.Println("Integer",v)
  case string:
    fmt.Println("string",v)
  default:
    fmt.Println("Unknow Type")
  }
}

func TestEmptyInterfaceAssertion(t *testing.T){
  DoSomething(10)
  DoSomething("hello")
}
