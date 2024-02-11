package main

import (
	"fmt"
	"games-email/services"
)

func main() {
	fmt.Println("Starting listener")
	services.StartListener()
}	