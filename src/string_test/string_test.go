package string_test

import "testing"
import "strings"
import "strconv"

func TestString(t *testing.T){
  var s string
  t.Log(s)
  s = "hello"
  t.Log(len(s))
  // s[1] = '3' //string是不可变的byte slice

  s = "中"
  t.Log(len(s)) //是byte数

  c := []rune(s)
  t.Log(len(c))
  t.Logf("中 unicode %x", c[0])
  t.Logf("中 UTF-8 %x", s)
}

func TestStringToRune(t *testing.T) {
  s := "中华人民共和国"
  for _, c := range s{
    t.Logf("%[1]c %[1]x", c)

  }
}

func TestStingFn(t *testing.T){
  s := "a,b,c"
  part := strings.Split(s, ",") //分割
  t.Log(part)
  t.Log(strings.Join(part, "-")) //合并

  s1 := strconv.Itoa(10) //整数转字符串
  t.Log("str" + s1)
}
