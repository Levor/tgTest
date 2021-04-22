package jobs

import (
	"io/ioutil"
	"log"
	"os"
)

func GetFiles() []string {
	files, err := ioutil.ReadDir(os.Getenv("FILES_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	var filesNames []string
	for _, file := range files {
		filesNames = append(filesNames, file.Name())

	}
	return filesNames
}
