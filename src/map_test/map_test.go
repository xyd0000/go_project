package map_test
import "testing"

func TestMapInit(t *testing.T) {
  m1 := map[int]int{1:1, 2:2, 3:3}  //[int]为key类型后面的是value类型
  t.Log(len(m1),m1[2])

  m2 := map[int]int{}
  m2[4] = 16
  t.Log(len(m2),m2)

  m3 := make(map[int]int, 10) //10代表的是容量，之所以不能指定长度是因为无法给切片初始化默认值
  t.Logf("len m3 = %d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
  m1 := map[int]int{}
  t.Log(m1[1]) //访问不存在的key，value默认为0, 不能返回nil
  m1[2] = 0
  t.Log(m1[2])
  if v, ok := m1[3]; ok{
    t.Logf("Key 3's value is %d", v)
  } else {
    t.Log("Key 3 is not exist")
  }
}

func TestTravelMap(t *testing.T){
  m1 := map[int]int{1:1, 2:2, 3:3}
  for k, v := range m1{
    t.Log(k, v)
  }
}
