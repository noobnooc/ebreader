package files

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

//Unepub Unzip the epub file
func Unepub(workingPath string, filePath string) error {
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return errors.New("Open file failed")
	}
	defer reader.Close()

	for _, file := range reader.File {
		lowerName := strings.ToLower(file.Name)
		if strings.HasPrefix(lowerName, "oebps") && !strings.HasSuffix(lowerName, "/") {
			newName := workingPath + file.Name[strings.Index(lowerName, "/"):]
			err := os.MkdirAll(getPath(newName), 0755)
			if err != nil {
				return err
			}

			writer, err := os.Create(newName)
			if err != nil {
				return err
			}
			defer writer.Close()

			reader, err := file.Open()
			if err != nil {
				return err
			}

			_, err = io.Copy(writer, reader)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//Clean Clear the working directory
func Clean(path string) error {
	if strings.HasSuffix(path, "ebreader") {
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func getPath(file string) string {
	path, _ := path.Split(file)
	return path
}
