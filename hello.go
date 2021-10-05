package main

import "fmt"

func main() {

	name := "Assis"
	version := 1.01

	fmt.Println("Hello", name, "This software is in", version, "version")

	fmt.Println("1 - Start monithoting")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")

	var command int
	fmt.Scan(&command)

	fmt.Println("The command chosen is", command)
}
