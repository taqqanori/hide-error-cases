package testinputs

func Literal() (int, error) {

	// error case inside lambda
	func() (int, error) {
		err := Error()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}()

	// error case inside function variable
	f := func() (int, error) {
		err := Error()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}
	f()

	return 1, nil
}
