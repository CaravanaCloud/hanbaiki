package cmd

import (
	"github.com/charmbracelet/log"
)

func Run(fn func() (interface{}, error), msg string) interface{} {
	result, err := fn()
	if err != nil {
		log.Fatal(msg, err)
	}
	return result
}
