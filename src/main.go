package main

import "fmt"
import "flag"

func main()  {
  fmt.Println("hello word")

  var name string
  flag.StringVar(&name,"name","everyone","The greeting object")

  age := flag.Int("age",18,"Input age")

  var married = flag.Bool("married",false,"Are you married")

  flag.Parse()
  fmt.Printf("hello %s!\n",name)
  fmt.Printf("%s'age is %d\n",name,*age)
  fmt.Printf("Married: %t\n",*married)
}
