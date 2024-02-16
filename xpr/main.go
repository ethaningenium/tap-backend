package main

import (
	"fmt"
	"tap/internal/libs/primitive"
)

func main() {
	obj1, err := primitive.GetObject("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj1)

	obj2, err := primitive.GetObject("65cf748159ff2a913ce8ba67")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj2)

	obj3, err := primitive.GetObject("ObjectId('65cf748159ff2a913ce8ba67')")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj3)

}

