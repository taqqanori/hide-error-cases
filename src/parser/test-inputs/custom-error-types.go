package testinputs

type HogeError struct{}

func CustomStructError() (int, *HogeError) {
	err := newHogeError()
	if err != nil {
		return -1, err
	}

	if err != nil {
		return -1, nil
	}

	return 1, nil
}

func newHogeError() *HogeError {
	return &HogeError{}
}

type HigeError interface{}

func CustomInterfaceError() (string, HigeError) {
	err := newHigeError()
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", nil
	}

	return "abc", nil
}

func newHigeError() HigeError {
	return ""
}
