package main

import (
	"fmt"

	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("@>@: ")
	log.SetFlags(0)

	names := []string{"zzjam", "xxjam", "yyjam"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	fmt.Println(quote.Go())
	fmt.Println("@>@")
}
