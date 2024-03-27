package utils

import (
	"fmt"
	"os"
)

func CreateFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func AppendToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	return nil
}
