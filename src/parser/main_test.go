package main

import (
	"testing"
)

func TestHoge(t *testing.T) {
	src := `package moke

func Hoge() (int, error) {
	func() {}()
	f := func() {}
	f()
	err := Hige()
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func Hige() error {
	return nil
}
`
	parse(src)
}

func Hoge() (int, error) {
	func() {}()
	f := func() {}
	f()
	err := Hige()
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func Hige() error {
	return nil
}
