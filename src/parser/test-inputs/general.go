package testinputs

import "fmt"

func General() (int, error) {

	// very general error case
	err := Error()
	if err != nil {
		fmt.Println("error")
		return -1, err
	}

	if err != nil {
		fmt.Println("error")
		return -1, nil
	}

	return 1, nil
}

func Error() error {
	return nil
}
