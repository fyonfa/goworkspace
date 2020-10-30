package main

import "fmt"

var a int

type hotdog int

var b hotdog

func hotDogType() {
	println("Hey this is a Fct: type hotdog")
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//a=b i can´t do this because they are diferent types, even hotdog is int type, b is type hotdog, not int
	//but we can do conversion because its underline type is integer int

	a = int(b) //this is call conversion, no casting like other languages
	fmt.Println("This is a converted")
	fmt.Println(a)

	fmt.Printf("%T", a)
}
