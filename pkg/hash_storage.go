package pkg

type Hash_Storage struct {
	files map[string]string
}

func (hash_storage Hash_Storage) Get_files() map[string]string {
	return hash_storage.files
}

func (hash_storage Hash_Storage) Get_bytes_from_file(key string) []byte {
	return []byte(hash_storage.files[key])
}

func New_Hash_Storage() *Hash_Storage {
	return &Hash_Storage{files: make(map[string]string)}
}
