package main

import (
	"fmt"
	"log"
	"strings"
	"task-manager-go/internal/app"
)

func main() {
	context := app.Context

	fmt.Println("Hello")
	for {
		fmt.Println("Please enter command:")
		input, err := context.Reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		err = context.CommandService.Run(strings.TrimSpace(input))
		if err != nil {
			log.Println(err)
		}
	}
}
