package file_sys

import "fmt"

func (fs *FileSystem) CreateDirectory(parent *Directory, dirName string, permissions string) *Directory {
	if _, exists := parent.SubDirs[dirName]; exists {
		fmt.Println("目录已存在！")
		return nil
	}

	dir := &Directory{
		Name:        dirName,
		Files:       make(map[string]*File),
		SubDirs:     make(map[string]*Directory),
		Permissions: permissions,
		Parent:      parent,
	}
	parent.SubDirs[dirName] = dir
	return dir
}

func (fs *FileSystem) DeleteDirectory(parent *Directory, dirName string) {
	dir, exists := parent.SubDirs[dirName]
	if !exists {
		fmt.Println("该目录不存在！")
		return
	}

	// 递归删除目录内容
	for _, subDir := range dir.SubDirs {
		fs.DeleteDirectory(dir, subDir.Name)
	}
	for fileName := range dir.Files {
		fs.DeleteFile(dir, fileName)
	}

	delete(parent.SubDirs, dirName)
}

func (fs *FileSystem) ListDirectory(dir *Directory) {
	for subDirName := range dir.SubDirs {
		fmt.Printf("目录 (%s) => %s\n", dir.Permissions, subDirName)
	}
	for fileName := range dir.Files {
		fmt.Printf("文件 (%s) => %s\n", dir.Permissions, fileName)
	}
}
