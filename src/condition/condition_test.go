package condition_test

import "testing"

func TestCondition_try(t *testing.T){
  if i:=1;i<3{
    t.Log("true")
  }
}


func TestSwitchMultiCase(t *testing.T){
  for i:=0;i<5;i++{
    switch i {
    case 0,2:
      t.Log("Even")
    case 1,3:
      t.Log("Old")
    default:
      t.Log("it is not 0-3")
    }
  }
}
