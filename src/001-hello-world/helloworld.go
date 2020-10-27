package main

import "fmt"

func main() {
	fmt.Println("hello everyone")
	foo()
	fmt.Println("hola")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

}
func foo() {
	fmt.Println("i am in foo")
}

func bar() {
	fmt.Println("yohooooo")
}

//control flow: (how a computer reads flow)
//(1) sequential or sequence
//(2) loop or iterative
//(3) conditional
