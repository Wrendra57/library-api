package compare

import (
	"fmt"
	"time"
)

func CompareTime(firstTime string, secondTime string) bool {

	// Parse the time strings, explicitly specifying their layouts and locations
	time1, err := time.Parse(time.RFC3339Nano, firstTime)
	if err != nil {
		fmt.Println("Error parsing first time:", err)
		return false
	}

	time2, err := time.Parse(time.RFC3339, secondTime) // Use time.RFC3339 for the second time without nanoseconds
	if err != nil {
		fmt.Println("Error parsing second time:", err)
		return false
	}

	// Compare the time values in UTC

	return time1.UTC().Before(time2.UTC())
}
