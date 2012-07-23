package main

import "errors"

func E(e error, s string) error {
	return errors.New(e.Error() + "\n" + s)
}