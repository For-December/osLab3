package file_sys

func NewFileSystem() *FileSystem {
	fs := &FileSystem{
		Root: &Directory{
			Name:        "/",
			Files:       make(map[string]*File),
			SubDirs:     make(map[string]*Directory),
			Permissions: "rwx",
		},
		FreeBlocks: make([]bool, TotalBlocks),
	}

	// 初始化所有块为空闲
	for i := 0; i < TotalBlocks; i++ {
		fs.FreeBlocks[i] = true
	}

	return fs
}
