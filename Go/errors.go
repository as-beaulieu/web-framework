package main

import (
	"fmt"
)

func errors(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
	}
}
