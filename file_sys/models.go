package file_sys

const (
	BlockSize   = 4096 // 每个Block大小
	TotalBlocks = 1024 // 总的Block数量
)

type File struct {
	Name        string
	Size        int
	Blocks      []int // 数据块指针
	Permissions string
}

type Directory struct {
	Name        string
	Files       map[string]*File
	SubDirs     map[string]*Directory
	Parent      *Directory
	Permissions string
}

type FileSystem struct {
	Root       *Directory
	FreeBlocks []bool // 空闲块管理
}

type Block struct {
	Data [BlockSize]byte
}
