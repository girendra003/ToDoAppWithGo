package cronjobs

import "fmt"

func Task1(value *int) func() {
	return func() {
		fmt.Println("Sending email to user ...", *value)
		*value++
	}
}
