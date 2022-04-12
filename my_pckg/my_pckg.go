package my_pckg

import "fmt"

type MyStruct struct {
	f1 string
}

func (m *MyStruct) SimpleMethod() {
	fmt.Println("This is SimpleMethod, current value of MyStruct->f1: ", m.f1)
}

func (m *MyStruct) ConcurrentMethod() {
	fmt.Println("This is ConcurrentMethod, current value of MyStruct->f1: ", m.f1)
}
