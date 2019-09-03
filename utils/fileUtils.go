package utils

import (
	"io/ioutil"
	"log"
	"os"
)

// ReadFromFile - reads from a file, delivers fallback, if reading not possible
func ReadFromFile(path string, fallback string) (string, error) {
	if path == "" {
		return fallback, nil
	}
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return fallback, err
	}
	return string(body), nil
}

// WriteToFile - writes content to a file
func WriteToFile(path string, content string) error {
	if path == "" {
		return nil
	}

	_, err := os.Create(path)
	if err != nil {
		log.Printf("unable to open %s", path)
		return err
	}

	data := []byte(content)
	err = ioutil.WriteFile(path, data, 0777)
	if err != nil {
		log.Printf("unable to write to %s", path)
		return err
	}
	return nil
}

// DeleteFile - removes a file
func DeleteFile(path string) error {
	var err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
