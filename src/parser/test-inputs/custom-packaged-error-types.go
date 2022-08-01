package testinputs

import errorpkg "github.com/taqqanori/toggle-error-cases/test-inputs/error-pkg"

func CustomPackagedStructError() (int, *errorpkg.PackagedStructError) {
	err := newStructError()
	if err != nil {
		return -1, err
	}

	if err != nil {
		return -1, nil
	}

	return 1, nil
}

func newStructError() *errorpkg.PackagedStructError {
	return &errorpkg.PackagedStructError{}
}

func CustomPackagedInterfaceError() (string, errorpkg.PackagedInterfaceError) {
	err := newHigeError()
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", nil
	}

	return "abc", nil
}

func newInterfaceError() errorpkg.PackagedInterfaceError {
	return ""
}
