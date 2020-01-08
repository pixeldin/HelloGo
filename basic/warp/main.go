package main

import (
	"fmt"
)

type Person struct {
	Age     int
	Name    string
	Friends []Person
}

func (*Person) Error() string { return "" }

func main() {
	defer fmt.Println("defer do something before panic...")
	fmt.Println("Ha haha...")
	panic("Something terrible!")
	// never mine defer handle
	//os.Exit(0)
}

func GetSomeWithWrapped() *Person {
	var p *Person
	wrong := false
	if wrong {
		return p
	} else {
		//return with type
		return nil
	}
}

func GetSomeNotNil() error {
	var p *Person = nil
	wrong := false
	if wrong {
		return p
	} else {
		//return p
		return p
	}
}


