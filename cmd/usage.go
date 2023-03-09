package cmd

import (
	"fmt"
)

func blue(s string) string {
	return s
}

func green(s string) string {
	return s
}

func rainbow(s string) string {
	return s
}

// rpadx adds padding to the right of a string.
func rpadx(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}
