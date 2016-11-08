//go:generate go-bind-plugin -plugin-package=./plugin -rebuild
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

	fmt.Printf("sin(0) = %f\n", plug.CalculateSin(100.00))
	plug.SayHello("Gophers!")

	fmt.Printf("Current year is: %d", plug.CurrentYear)
	fmt.Println(plug)
}
