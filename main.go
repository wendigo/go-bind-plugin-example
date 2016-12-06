//go:generate go-bind-plugin -plugin-package=./plugin -rebuild -dereference-vars
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("> plug, _ := BindPluginAPI(\"./plugin.so\")")
	plug, err := BindPluginAPI("./plugin.so")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("> plug.CalculateSin(100.0)")
	fmt.Println(plug.CalculateSin(100.00))
	fmt.Println("> plug.SayHello(\"Gophers!\")")
	plug.SayHello("Gophers!")

	fmt.Println("> plug.CurrentYear")
	fmt.Println(plug.CurrentYear)
	
	fmt.Println("> plug")
	
	fmt.Println(plug)
}
