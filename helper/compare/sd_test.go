package compare

import (
	"fmt"
	"testing"
)

func TestAuthenticateSucces(t *testing.T) {
	firstTime := "2024-02-06T18:33:08.060151463+07:00"
	secondTime := "2024-02-07T14:12:59Z"

	if CompareTime(firstTime, secondTime) {
		fmt.Println("The first time is before the second time.")
	} else {
		fmt.Println("The first time is not before the second time.")
	}
}
