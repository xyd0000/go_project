package arry_test
import "testing"

func TestArryInit(t *testing.T) {
  var arr [3]int
  arr1 := [4]int{1,2,3,4}
  arr2 := [...]int{1,2,3,4,5}
  arr1[1] = 5
  t.Log(arr[1], arr[2])
  t.Log(arr1, arr2)
}

func TestArryTravel(t *testing.T){
  arr := [...]int{1,2,3,4,5}
  for i := 0; i < len(arr); i++ {
    t.Log(arr[i])
  }
  for idx, _ := range arr{
    t.Log(arr[idx])
  }
}

func TestArrySection(t *testing.T){
  arr := [...]int{1,2,3,4,5}
  arr1_sec := arr[:3]
  t.Log(arr1_sec)
}

func TestArry_try(t *testing.T){
  a := [3]int{1,2,3} //声明和赋值
  t.Log(a)

  for _ , e :=range a {//_ 是返回值站位
    t.Log(e)
  }

  b := [2][2]int{{1,2},{3,4}} //多维数组
  t.Log(b)

  d := [...]int{1,2,3,4,5} //初始化时用...代表长度
  t.Logf("%T",d)
  //append(d,6) //对数组进行扩容会报错，即数组是定长的
}
