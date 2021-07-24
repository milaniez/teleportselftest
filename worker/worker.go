package worker

import "fmt"

func Hello(name string) string {
	ret := fmt.Sprintf("Hi %v", name)
	return ret
}