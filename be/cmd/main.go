package main

import (
	"expense-management-system/internal/bootstrap"
	"log"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
