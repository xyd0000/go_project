package type_test

import "testing"
/*
1.不支持隐式类型转化，只支持显示类型转化
*/
type MyInt int64
func TestType(t *testing.T){
  var a int32 = 1
  var b int64
  b = int64(a)
  var c MyInt
  //c = b //不支持别名类型的转化
  c = MyInt(b)
  t.Log(a ,b ,c)
}

func TestPoint(t *testing.T){
  a :=1
  aPtr := &a
  //aPtr = aPtr + 1 //不支持指针运算
  t.Log(a,aPtr)
  t.Logf("%T, %T",a,aPtr)
}


func TestString(t *testing.T){
  var s string
  t.Log("*" + s + "*")
  t.Log(len(s))
}
