package main

import (
	"fmt"
)
import "AwesomeProject/item"

//go:generate mockgen -source=main.go -destination=./main_mock.go -package=main

type MyStructer interface {
	CreateMyStruct() MyStruct
}

type MyStruct struct {
	a string
}

func (m MyStruct) CreateMyStruct() MyStruct {
	return MyStruct{
		a: "some value",
	}
}

func main() {
	//PublicPackage.SayMe()
	//simply_another_package.SayMe()
	//
	//items := item.NewItems()
	//
	//items.Add(item.NewItem(1, "Some item"))
	//items.Add(item.NewItem(3, "Some item2"))
	//IterateThroughItems(items)
	//
	//aItems := item.NewAnotherItems()
	//aItems.AddAnotherItem(item.NewAnotherItem(7, "Another item"))
	//IterateThroughAnotherItems(aItems)
	//
	//c := Config{
	//	FirstField:  "123",
	//	secondField: "234",
	//}
	//fmt.Println(c)
}

func IterateThroughAnotherItems(ais item.AnotherItems) bool {
	for _, i := range ais.AnotherElements() {
		fmt.Println("ITERATION ANOTHER ITEM", i)
	}
	return true
}

func IterateThroughItems(is *item.Isnterface) bool {
	for _, i := range (*is).Elements() {
		fmt.Println("ITERATION ", *(&i))
		s, e := i.GetNameWithValidation()
		if e != nil {
			fmt.Println(e)
			return false
		}
		fmt.Printf("VALIDATED: %s", s)
	}
	return true
}
