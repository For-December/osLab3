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
		data = append(data, fs.GetBlockData(blockIndex)...)
	}
	return data
}

// GetBlockData 模拟获取块数据
func (fs *FileSystem) GetBlockData(blockIndex int) []byte {
	// 从模拟原始块（磁盘）中获取数据
	return fs.MockRawBlocks[blockIndex].Data[:]
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
		// 写入数据到块，这里限定了每个块的大小为 BlockSize
		fs.WriteBlockData(blockIndex, data[i*BlockSize:min((i+1)*BlockSize, len(data))])
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

// WriteBlockData 模拟写入块数据
func (fs *FileSystem) WriteBlockData(blockIndex int, data []byte) {
	// 写入到模拟原始块（磁盘）
	// 已限定了每个块的大小为 BlockSize
	copy(fs.MockRawBlocks[blockIndex].Data[:], data)
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

func (fs *FileSystem) ReadFileContent(dir *Directory, fileName string) []byte {
	file, exists := dir.Files[fileName]
	if !exists {
		fmt.Println("File does not exist")
		return nil
	}

	return fs.ReadFile(file)
}

func (fs *FileSystem) AppendToFile(dir *Directory, fileName string, data []byte) {
	file, exists := dir.Files[fileName]
	if !exists {
		fmt.Println("File does not exist")
		return
	}

	existingData := fs.ReadFile(file)
	newData := append(existingData, data...)
	fs.WriteFile(file, newData)
}
