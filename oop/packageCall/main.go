package main

import (
	"fmt"
	"reflect"

	"github.com/golang/oop/packageCall/pack"
)

func main() {

	P := pack.Protocols{}

	fmt.Println(reflect.TypeOf(P))

	P.Fircall()
}
