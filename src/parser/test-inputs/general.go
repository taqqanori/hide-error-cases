package testinputs

func Hoge() (int, error) {

	// very general error case
	err := Hige()
	if err != nil {
		return -1, err
	}

	// error case inside lambda
	func() (int, error) {
		err := Hige()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}()

	// error case inside function variable
	f := func() (int, error) {
		err := Hige()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}
	f()

	return 1, nil
}

// named return type cases
func Huge() (ret int, err error) {

	// very general error case
	err = Hige()
	if err != nil {
		return -1, err
	}

	// error case inside lambda
	func() (ret int, err error) {
		err = Hige()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}()

	// error case inside function variable
	f := func() (ret int, err error) {
		err = Hige()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}
	f()

	return 1, nil
}

func Hige() error {
	return nil
}
