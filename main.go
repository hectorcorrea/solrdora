package main

import (
	"gosiah/web"
	// "log"
	// "reflect"
)

func main() {

	// var s interface{}
	// s = "hello"
	//
	// var a []interface{}
	// a = append(a, "hello")
	// a = append(a, "world")
	//
	// log.Printf("%T", a)
	// log.Printf("%T", s)
	//
	// log.Printf("typeOf() %s", reflect.TypeOf(a))
	// log.Printf("kind: %s", reflect.TypeOf(a).Kind())
	// log.Printf("slice? %s", reflect.TypeOf(a).Kind() == reflect.Slice)

	// t := reflect.ValueOf(s).Kind()
	// log.Printf("%s", t)
	// log.Printf("%v", t == "slice")

	// return

	address := "localhost:9001"
	web.StartWebServer(address)
}
