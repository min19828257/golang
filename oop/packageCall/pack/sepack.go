package pack

import (
	"fmt"
)

type Protocols struct {
}

func (p Protocols) secondcall() {
	fmt.Println("secondCall")
}
