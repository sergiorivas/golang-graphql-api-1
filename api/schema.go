package api

import (
	"io/ioutil"
)

func GetSchema(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
