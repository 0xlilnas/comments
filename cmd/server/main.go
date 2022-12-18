package main

import "fmt"

func Run() error {
	fmt.Println("Starting our application")
	return nil
}

func main() {
	fmt.Println("Hello Commenting System!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
