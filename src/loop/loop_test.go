package loop_test

import "testing"

func TestLoop_try(t *testing.T){
  n := 0
  for n<5 {
    n++
    t.Log(n)
  }
}
