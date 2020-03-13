package lib

import (
	"fmt"
	"time"
)

// Timer prints how long a routine takes
func Timer(start time.Time, message string) {
	time := time.Since(start)
	fmt.Printf("\n%v: %v\n", message, time)
}
