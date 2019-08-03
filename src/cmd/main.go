package main

import (
	"handler"
	"net/http"
	"fmt"

	"cmd/flags"
)

func main() {
	fmt.Println("Port Number:", flags.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", flags.Port), handler.Router)
	if err != nil {
		fmt.Print(err)
	}	
}
