package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func valid_file(file string, root string) bool {
	abs_file, _ := filepath.Abs(file)
	abs_root, _ := filepath.Abs(root)
	return !(strings.HasSuffix(abs_file, ".ags") && strings.EqualFold(abs_file, abs_root))
}

func Scan_folder(root string) Hash_Storage {
	fmt.Printf("Argus is scanning: %s\n", root)

	hash_storage := New_Hash_Storage()

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && valid_file(path, root) {
			file_bytes, _ := os.ReadFile(path)
			file_hash := hasher(string(file_bytes))
			hash_storage.files[path] = file_hash
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return *hash_storage
}
