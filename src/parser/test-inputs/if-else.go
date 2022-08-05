package testinputs

import "fmt"

func Else() (int, error) {
	err := Error()
	if err == nil {
		return 1, nil
	} else {
		fmt.Println("error")
		return -1, err
	}
}

func ElseNoError() (int, error) {
	err := Error()
	if err != nil {
		fmt.Println("error")
		return -1, err
	} else {
		return -1, nil
	}
}

func ElseIf() (int, error) {
	err := Error()
	if err == nil {
		return 1, nil
	} else if err != nil {
		fmt.Println("error")
		return -1, err
	}
	return 1, nil
}

func ElseIfNoError() (int, error) {
	err := Error()
	if err != nil {
		fmt.Println("error")
		return -1, err
	} else if err == nil {
		return 1, nil
	}
	return 1, nil
}

func MultiElses1() (int, error) {
	err := Error()
	if err != nil {
		fmt.Println("error")
		return -1, err
	} else if err == nil {
		return 1, nil
	} else {
		return 1, nil
	}
}

func MultiElses2() (int, error) {
	err := Error()
	if err == nil {
		return 1, nil
	} else if err != nil {
		fmt.Println("error")
		return -1, err
	} else {
		return 1, nil
	}
}

func MultiElses3() (int, error) {
	err := Error()
	if err == nil {
		return 1, nil
	} else if err != nil {
		return 1, nil
	} else {
		fmt.Println("error")
		return -1, err
	}
}

func NestedIfElse() (int, error) {
	err := Error()
	if err == nil {
		if err == nil {
			fmt.Println("abc")
		} else if err != nil {
			return -1, err
		} else {
			return -1, err
		}
		return -1, err
	} else if err != nil {
		if err == nil {
			return -1, err
		} else if err != nil {
			fmt.Println("abc")
		} else {
			return -1, err
		}
		return 1, nil
	} else {
		if err == nil {
			return -1, err
		} else if err != nil {
			return -1, err
		} else {
			fmt.Println("abc")
		}
		return -1, err
	}
}

func LineBreakAfterElse() (int, error) {
	err := Error()
	if err == nil {
		return 1, nil
	} else
	{
		fmt.Println("error")
		return -1, err
	}
}

