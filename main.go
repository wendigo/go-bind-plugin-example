//go:generate go-bind-plugin -plugin-package=./plugin -rebuild -dereference-vars
package main

import (
	"fmt"
	"log"
)

func main() {
	plug, err := BindPluginAPI("./plugin.so")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("plug.CalculateSin(100.0) = %f\n", plug.CalculateSin(100.00))
	plug.SayHello("Gophers!")

	fmt.Printf("plug.CurrentYear is: %d\n\n\n", plug.CurrentYear)

	fmt.Println(plug)
}
