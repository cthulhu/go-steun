package fixture

import (
	"io/ioutil"
	"os"
)

func Read(filename string) []byte {
	file, err := os.Open("fixtures/" + filename)
	if err != nil {
		panic(err)
	}

	d, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return d
}
