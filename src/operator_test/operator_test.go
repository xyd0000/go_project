package operator_test

import "testing"

func TestCompareArry(t *testing.T){
  a:=[...]int{1,2,3,4}
  b:=[...]int{1,2,4,5}
  //c:=[...]int{1,2,3,4,5}
  d:=[...]int{1,2,3,4}
  t.Log(a == b)
  //t.Log(a == c) //不支持长度不同的数组进行比较
  t.Log(a == d)
}
