package lib

import (
	"fmt"
	"strings"
	"time"
)

// Timer prints how long a routine takes
func Timer(start time.Time, message string) {
	time := time.Since(start)
	fmt.Printf("\n%v: %v\n", message, time)
}

func convertStringSliceToString(slice []string) string {
	return strings.Join(slice, "\n")
}
