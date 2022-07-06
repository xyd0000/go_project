package func_test
import "testing"
import "time"
import "fmt"
import "math/rand"

func returnMultiValues() (int,int){
  return rand.Intn(10), rand.Intn(20)
}


func calculateTimeSpend(inner func(op int) int) func(op int) int{
  return func(n int) int {
    start := time.Now()
    ret := inner(n)
    fmt.Println("time spent:", time.Since(start).Seconds())
    return ret
  }
}

func slowFun(op int)  int {
  time.Sleep(time.Second*1)
  return op
}

func TestFn(t *testing.T){
  a, _ := returnMultiValues()
  t.Log(a)

  tsSf := calculateTimeSpend(slowFun)
  t.Log(tsSf(10))
}


func Sum(ops ...int) int { //可变长参数
  sum := 0
  for _, value := range ops{
    sum += value
  }
  return sum
}

func TestSum(t *testing.T) {
  t.Log(Sum(1,2,3))
  t.Log(Sum(1,2,3,4))
}


func Clear() {
  fmt.Println("clear resources")
}
func TestDefer(t *testing.T) {
  defer Clear()  //在运行结束执行
  fmt.Println("testdefer")
  panic("err")
}
