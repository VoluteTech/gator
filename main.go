package main

import (
	"fmt"
	"log"

	"github.com/VoluteTech/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config: %v", config)
	config.SetUser("volute")
	fmt.Printf("Config after setting username: %v", config)
}
