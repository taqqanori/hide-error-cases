package testinputs

import "fmt"

func Loops() (string, error) {

	for i := range []int{1, 2, 3} {
		if i == 4 {
			fmt.Println("error")
			return "", fmt.Errorf("error")
		}
	}

	return "hoge", nil
}
