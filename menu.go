package main

import "fmt"

func main() {
    var str string

	fmt.Println("Menu Program:")
	for{
		fmt.Scan(&str)
		switch  str {
		case "hello":
			fmt.Println("world")
		case "q":
			return
		}
	}
}