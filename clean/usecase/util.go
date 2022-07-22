package usecase

import "io/ioutil"

func IsCommonFileExist(dirPath string, fileName string) bool {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.Name() == fileName {
			return true
		}
	}

	return false
}
