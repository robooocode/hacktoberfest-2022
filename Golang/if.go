
package main

import "fmt"

func main() {
	var name string = "jokowidodo"

	if name == "reza" {
		fmt.Println("hello " + name)
	} else if name == "joko" {
		fmt.Println("terima kasih joko")
	} else {
		fmt.Println("your name not have here")
	}

	//short statetment cuman ada di golang
	if lenght := len(name); lenght > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("name sudah benar")
	