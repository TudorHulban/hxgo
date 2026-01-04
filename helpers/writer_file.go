package helpers

import "os"

func GetFileWriter(filename string) (*os.File, error) {
	file, errCreate := os.Create(filename)
	if errCreate != nil {
		return nil,
			errCreate
	}

	return file,
		nil
}
