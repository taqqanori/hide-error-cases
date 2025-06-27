package testinputs

func NamedReturnValues() (ret int, err error) {

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

func NamedReturnValuesMultipleForSingleType() (i1, i2 *int, err error) {
	err = Error()
	if err != nil {
		return i1, nil, err
	}
	return i1, i2, nil
}

func NamedReturnValuesWithReturnWithoutArgs1() (err error) {
	err = Error()
	something := 1

	if err != nil {
		panic(err)
	}

	if err != nil {
		return
	}

	if nil != err {
		return
	}

	if err != nil && 0 < something {
		return
	}

	if err != nil || 0 < something {
		return
	}

	if 0 < something && err != nil {
		return
	}

	if 0 < something || err != nil {
		return
	}

	return
}

func NamedReturnValuesWithReturnWithoutArgs2() (i int, err1, err2 error) {
	err1 = Error()
	something := 1

	if err1 != nil {
		return
	}

	if err2 != nil {
		return
	}

	if err1 != nil && err2 != nil {
		return
	}

	if err1 != nil || err2 != nil {
		return
	}

	if err1 != nil && 0 < something {
		return
	}

	if err1 != nil || 0 < something {
		return
	}

	if err1 != nil && 0 < something && something < 1 {
		return
	}

	if err1 != nil && 0 < something || something < 1 {
		return
	}

	if err1 != nil || 0 < something && something < 1 {
		return
	}

	if err1 != nil && 0 < something || something < 1 {
		return
	}

	if err1 != nil && (0 < something && something < 1) {
		return
	}

	if err1 != nil && (0 < something || something < 1) {
		return
	}

	if err1 != nil || (0 < something && something < 1) {
		return
	}

	if err1 != nil || (0 < something || something < 1) {
		return
	}

	if 0 < something && (err1 != nil && something < 1) {
		return
	}

	if 0 < something && (err1 != nil || something < 1) {
		return
	}

	if 0 < something || (err1 != nil && something < 1) {
		return
	}

	if 0 < something || (err1 != nil || something < 1) {
		return
	}

	if err1 != nil && err2 != nil && 0 < something {
		return
	}

	if err1 != nil && err2 != nil || 0 < something {
		return
	}

	if err1 != nil || err2 != nil && 0 < something {
		return
	}

	if err1 != nil || err2 != nil || 0 < something {
		return
	}

	if err1 != nil && (err2 != nil || 0 < something) {
		return
	}

	if (err1 != nil || err2 != nil) && 0 < something {
		return
	}

	if (err1 != nil || err2 != nil) || 0 < something {
		return
	}

	return
}
