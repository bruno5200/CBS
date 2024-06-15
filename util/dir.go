package util

import "os"

// Check if directory exists and create it if not
func CheckDir(dir string) error {

	// Check directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {

		// Create directory if not exists
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
