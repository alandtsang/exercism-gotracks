// Package twofer implements given a name, return a string with the message:
// One for X, one for me.
package twofer

import "fmt"

// ShareWith returns different message according to name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
