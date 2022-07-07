package testinputs

func IfOutsideFunc() (int, error) {
	err := Error()
	if err != nil {
		func() (string, error) {
			return "", err
		}()
		f := func() (string, error) {
			return "", err
		}
		f()
	}
	return 1, nil
}
