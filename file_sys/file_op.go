package file_sys

import "fmt"

func (fs *FileSystem) CreateFile(dir *Directory, fileName string, permissions string) *File {
	if _, exists := dir.Files[fileName]; exists {
		fmt.Println("File already exists")
		return nil
	}

	file := &File{
		Name:        fileName,
		Size:        0,
		Blocks:      make([]int, 0),
		Permissions: permissions,
	}
	dir.Files[fileName] = file
	return file
}

func (fs *FileSystem) ReadFile(file *File) []byte {
	var data []byte
	for _, blockIndex := range file.Blocks {
		// 读取每个块的数据
		data = append(data, GetBlockData(blockIndex)...)
	}
	return data
}

// 模拟获取块数据
func GetBlockData(blockIndex int) []byte {
	// 在真实实现中，这里会从磁盘读取数据
	return []byte{} // 这里返回空数据
}

func (fs *FileSystem) WriteFile(file *File, data []byte) {
	// 清空文件现有数据块
	file.Blocks = file.Blocks[:0]

	// 计算需要多少块
	numBlocks := (len(data) + BlockSize - 1) / BlockSize

	for i := 0; i < numBlocks; i++ {
		// 分配新的块
		blockIndex := fs.allocateBlock()
		if blockIndex == -1 {
			fmt.Println("No free blocks available")
			return
		}
		file.Blocks = append(file.Blocks, blockIndex)
		// 写入数据到块
		WriteBlockData(blockIndex, data[i*BlockSize:min((i+1)*BlockSize, len(data))])
	}

	file.Size = len(data)
}

// 分配空闲块
func (fs *FileSystem) allocateBlock() int {
	for i := 0; i < TotalBlocks; i++ {
		if fs.FreeBlocks[i] {
			fs.FreeBlocks[i] = false
			return i
		}
	}
	return -1
}

// 模拟写入块数据
func WriteBlockData(blockIndex int, data []byte) {
	// 在真实实现中，这里会写入数据到磁盘
}

func (fs *FileSystem) DeleteFile(dir *Directory, fileName string) {
	file, exists := dir.Files[fileName]
	if !exists {
		fmt.Println("File does not exist")
		return
	}

	// 释放文件占用的块
	for _, blockIndex := range file.Blocks {
		fs.FreeBlocks[blockIndex] = true
	}

	delete(dir.Files, fileName)
}
