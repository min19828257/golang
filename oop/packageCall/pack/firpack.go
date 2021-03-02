package pack

import (
	"fmt"
	"reflect"
)

func (p Protocols) Fircall() {
	fmt.Println("firstCall", reflect.TypeOf(p), p)
	p.secondcall()
}
