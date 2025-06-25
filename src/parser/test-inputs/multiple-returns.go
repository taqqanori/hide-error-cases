package testinputs

import "os"

// https://github.com/taqqanori/hide-error-cases/issues/6
func Func() (*os.File, error) {
	if file, err := os.Open("data.txt"); err == nil {
		switch err = file.Chmod(os.ModeAppend); err {
		default:
			return nil, err
		case nil:
			return file, nil
		}
	} else {
		return nil, err
	}
}
