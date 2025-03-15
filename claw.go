package claw

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func BundleWebComponents(dir string, out string) error {

	outFile := ""
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		pathParts := strings.Split(path, ".")
		ext := pathParts[len(pathParts)-1]
		if ext == "js" {
			bytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			file := string(bytes)
			words := strings.Split(file, " ")
			if len(words) < 2 {
				return fmt.Errorf(`component at [%s] is malformed, it has less than 2 words`, path)
			}
			className := words[1]
			webComponentName, err := convertToKebabCaseStrict(className)
			if err != nil {
				return err
			}
			defineStatement := fmt.Sprintf(`customElements.define("%s", %s)`, webComponentName, className)
			outFile = outFile + file + "\n\n" + defineStatement + "\n\n"
		}
		return nil
	})
	if err != nil {
		return err
	}
	outFile = strings.TrimSpace(outFile)

	err = writeToFile(out, outFile, true)
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(path string, content string, overwrite bool) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}
	if _, err := os.Stat(path); err == nil {
		if !overwrite {
			return errors.New("file already exists and overwrite is false")
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check file existence: %w", err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

func convertToKebabCaseStrict(input string) (string, error) {
	re := regexp.MustCompile(`[A-Z][a-z]*`)
	parts := re.FindAllString(input, -1)
	if len(parts) != 2 {
		return "", errors.New("input must contain exactly two words in PascalCase")
	}
	return strings.ToLower(parts[0]) + "-" + strings.ToLower(parts[1]), nil
}
