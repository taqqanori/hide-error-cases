package testinputs

func Nested() (int, error) {
	err := Error()

	if err != nil {
		if err != nil {
			return -1, err
		}
		return -1, err
	}

	func() {
		f := func() (int, error) {
			err := Error()
			if err != nil {
				return -1, err
			}
			return 1, nil
		}
		f()
	}()

	return 1, nil
}
