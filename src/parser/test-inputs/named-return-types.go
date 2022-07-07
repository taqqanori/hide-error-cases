package testinputs

func NamedReturnTypes() (ret int, err error) {

	// very general error case
	err = Error()
	if err != nil {
		return -1, err
	}

	// error case inside lambda
	func() (ret int, err error) {
		err = Error()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}()

	// error case inside function variable
	f := func() (ret int, err error) {
		err = Error()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}
	f()

	return 1, nil
}
