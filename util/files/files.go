package files

import (
	"archive/zip"
	"ebreader/config"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

var ()

//Unepub Unzip the epub file
func Unepub() error {
	reader, err := zip.OpenReader(config.File)
	if err != nil {
		return errors.New("Open file failed")
	}
	defer reader.Close()

	for _, file := range reader.File {
		lowerName := strings.ToLower(file.Name)
		if strings.HasPrefix(lowerName, "oebps") && !strings.HasSuffix(lowerName, "/") {
			newName := config.Path + file.Name[strings.Index(lowerName, "/"):]
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
func Clean() error {
	if strings.HasSuffix(config.Path, "ebreader") {
		err := os.RemoveAll(config.Path)
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
