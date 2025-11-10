package main

import "github.com/valentinusdelvin/savebite-be/internal/bootstrap"

func main() {
	err := bootstrap.Start()
	if err != nil {
		panic(err)
	}
}
