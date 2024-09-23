package timer

import (
	"fmt"
	"time"
)

func PrintTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
