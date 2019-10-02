package schema

import "io/ioutil"

func GetSchema() (string, error) {
	b, err := ioutil.ReadFile("./general.graphql")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
