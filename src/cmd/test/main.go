package main

import (
	//"github.com/mojosa-software/godat/src/mapx"
	//"github.com/mojosa-software/godat/src/slicex"
	//"github.com/mojosa-software/godat/src/poolx"
	"github.com/mojosa-software/godat/src/rangex"
	"github.com/mojosa-software/godat/src/mapx"
	"fmt"
)

type Struct struct {
	Name string
	Value int
}

type MyMap struct {
	mapx.Map[string, int]
}

func NewMyMap() *MyMap {
	return &MyMap{
		Map: mapx.New[string, int](),
	}
}

func main() {
	rangex.New[float32](0, .001, 0.050).Chan().ForEach(
		func(i int, v float32) bool {
			fmt.Println(i, v)
			return true
		},
	)
	
	m := mapx.New[string, int]()
	m.Set("suck", 1)
	m.Set("cock", 10)
	
	for k, v := range m {
		fmt.Println(k, v)
	}	
	
	fmt.Println(m.Has("dick"))
	
	mm := NewMyMap()
	mm.Set("dicker", 15)
	
	fmt.Println(mm.Get("dicker"))
	/*m := map[string] string {
		"Key1" : "Value1",
		"Key2" : "Value2",
		"Key3" : "Value3",
	}
	m1 := map[int] string {
		1 : "Val1",
		2 : "Val2",
		7 : "Val7",
	}
	s := []Struct {
		{"Name1", 1},
		{"Name2", 2},
	}

	fmt.Println(m)
	fmt.Println(slicex.MakeMap(
		s,
		func(s []Struct, i int) string {
			return s[i].Name
		},
	))

	fmt.Printf("%q\n", mapx.Keys(m))
	fmt.Printf("%q\n", mapx.Values(m))
	fmt.Printf("%q\n", mapx.Reverse(m))
	fmt.Printf("%v\n", mapx.Reverse(m1))
	
	ll := poolx.New[int]()
	ll.Append(0)
	ll.Append(1)
	ll.Append(2)
	ll.Del(256)
	ll.Del(1)
	for p := range ll.Range() {
		fmt.Println(p)
	}*/
}
