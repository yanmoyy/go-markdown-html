package gen

import (
	"fmt"
	"os"
)

func CopyStaticFilesRecursive(src, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		err = os.MkdirAll(dest, 0750)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	dir, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}
	for _, file := range dir {
		from := src + "/" + file.Name()
		to := dest + "/" + file.Name()
		fmt.Println("*", from, "->", to)
		if file.IsDir() {
			err := CopyStaticFilesRecursive(from, to)
			if err != nil {
				return err
			}
		} else {
			// #nosec: G304
			data, err := os.ReadFile(from)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}
			err = os.WriteFile(to, data, 0600)
			if err != nil {
				return fmt.Errorf("failed to write file: %w", err)
			}
		}
	}
	return nil
}
