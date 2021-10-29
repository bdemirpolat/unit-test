package main

import (
	"fmt"
)

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hello Anonymous!"
	}

	if len(name) >= 5 {
		name = fmt.Sprintf("%v...", name[:5])
	}

	return fmt.Sprintf("Hello %s!", name)
}
