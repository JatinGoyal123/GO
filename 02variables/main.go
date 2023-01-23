package main

import "fmt"

const LoginToken int = 122 // public
func main() {
	var username string = "jatin"
	fmt.Println(username)
	fmt.Printf("The type of var is %T\n", username)

	var isLogin bool = false
	fmt.Println(isLogin)
	fmt.Printf("The type of var is %T\n", isLogin)

	// inline
	var name = "Sunil"
	fmt.Println(name)
	fmt.Printf("The type of var is %T\n", name)

	// using walrus operator
	d := 255.55
	fmt.Println(d)
	fmt.Printf("The type of var is %T\n", d)

	fmt.Println(LoginToken)
}

//fp for short of fmt.Println("hello")
