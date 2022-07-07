package testinputs

func MultipleErrors() (error, string, error, int, error) {
	err := Error()

	if err != nil {
		return err, "", nil, -1, nil
	}

	if err != nil {
		return nil, "", err, -1, nil
	}

	if err != nil {
		return nil, "", nil, -1, err
	}

	if err != nil {
		return err, "", nil, -1, err
	}

	if err != nil {
		return err, "", err, -1, err
	}

	if err != nil {
		return nil, "", nil, -1, nil
	}

	return nil, "hoge", nil, 1, nil
}
