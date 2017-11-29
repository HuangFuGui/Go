package main

import (
	"fmt"
	"reflect"
)

type PreResponse struct {

}

type Response struct {
	PreResponse
}

func (pre *PreResponse) BaseShow(){
	fmt.Printf("BaseShow()\n")
}

func (pre *PreResponse) basewatch(){
	fmt.Printf("basewatch()\n")
}

func (resp *Response) Show(){
	fmt.Printf("Hello Show()\n")
}

func (resp *Response) watch(){
	fmt.Printf("Hello watch()\n")
}

var m map[string]interface{} = make(map[string]interface{})

func main(){
	//反射方式一：只能大写方法
	s := "resp"
	m[s] = &Response{}
	fmt.Printf("%v\n", m["resp"])
	value := reflect.ValueOf(m[s])
	fmt.Printf("name:%v\n", value.Type().Elem().Name())
	value.MethodByName("BaseShow").Call([]reflect.Value{})
	value.MethodByName("Show").Call([]reflect.Value{})

	/*输出：
	&{{}}
	name:Response
	BaseShow()
	Hello Show()
	*/

	//反射方式二：根据接口反射，可以反射调用小写方法，参照fly
}
