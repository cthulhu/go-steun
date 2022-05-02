package fixture

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Read(filename string) []byte {
	return ReadFile("fixtures/" + filename)
}

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	d, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return d
}

func Path(fixtureName string) string {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(currentPath, "fixtures", fixtureName)
}
