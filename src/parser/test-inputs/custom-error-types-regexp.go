package testinputs

type HogeException struct{}

func CustomStructException() (int, *HogeException) {
	err := newHogeException()
	if err != nil {
		return -1, err
	}

	if err != nil {
		return -1, nil
	}

	return 1, nil
}

func newHogeException() *HogeException {
	return &HogeException{}
}

type HigeException interface{}

func CustomInterfaceException() (string, HigeException) {
	err := newHigeException()
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", nil
	}

	return "abc", nil
}

func newHigeException() HigeException {
	return ""
}
