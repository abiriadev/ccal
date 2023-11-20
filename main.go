package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Current time: %s\n", t.Local())
}
