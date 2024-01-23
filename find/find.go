package find

import (
	"crypto/sha256"
	"io"
	"os"
	"path/filepath"
)

func hashfile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// FindDuplicates finds duplicate files in a directory.
// TODO: Have this return a map[string][]string where the key is a filename and the value is a list of files that duplicate it.
func FindDuplicates(dir string) (dupes []string, err error) {
	hashes := make(map[string]string)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			hash, err := hashfile(path)
			if err != nil {
				return err
			}

			_, ok := hashes[string(hash)]
			if ok {
				dupes = append(dupes, path)
				return nil
			}

			hashes[string(hash)] = path
		}
		return nil
	})

	return
}
