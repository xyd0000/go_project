package arryAndslice
import "fmt"
import "testing"

func TestSliceInit(t *testing.T){
  var s0 []int
  t.Log(len(s0), cap(s0))
  s0 = append(s0,1)
  t.Log(len(s0), cap(s0))

  s1 := []int{1,2,3,4}
  t.Log(len(s1), cap(s1))
}

func printSlice(name string,x []int){
  fmt.Printf("%s: len = %d, cap = %d, value= %v\n",name ,len(x),cap(x),x)

}

func TestSlice_try(t *testing.T){
  //未初始化切片容量时，长度和容量一样
  s1 := make([]string, 5)
  t.Log(len(s1), cap(s1), s1)

  //长度和容量可以不一样
  s2 := make([]int, 5, 6)
  t.Log(len(s2), cap(s2), s2)

  //fmt.Printf("The content of s2: %s\n",s2[5]) 不可以访问在容量内但是不在长度外的数据
  s2 = append(s2,1)
  printSlice("After append ,s2",s2)
  s2 = append(s2,2)
  printSlice("After append ,s2",s2)
  fmt.Println()

  s3 := []int{1,2,3,4,5,6,7,8}
  t.Logf("%T",s3)
  s4 := s3[3:6]
  printSlice("s4",s4)
  s4 = s4[0:cap(s4)] //扩大切片的窗口
  printSlice("s4",s4)
}

func TestSliceGrowing(t *testing.T) {
  s := []int{}
  for i := 1 ; i < 9 ; i++ {
    s = append(s, i)
    t.Log(len(s), cap(s), s)
  }
}

func TestSliceShareMemory(t *testing.T){
  year := []string{"jan", "Feb", "Mar", "Apr", "May", "jun", "jul", "Aug", "sep",
  "Oct", "Nov", "Dec"}
  Q2 := year[3:6]
  t.Log(len(Q2), cap(Q2), Q2)

  summer := year[5:8]
  t.Log(len(summer), cap(summer), summer)
  summer[0] = "Unknow"
  t.Log(Q2)
  t.Log(year)
}

func TestSliceComparing(t *testing.T) {
  a := []int{}
  b := []int{}
  if a == b{
    t.Log("compare")
  }
}
