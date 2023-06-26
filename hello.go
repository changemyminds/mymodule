package mymodule

import "fmt"

// Hello returns a greeting
func Hello(name string) string {
	return fmt.Sprintf("Hello, demo %s!", name)
}
