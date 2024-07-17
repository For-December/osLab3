package main

import (
	"bufio"
	"fmt"
	"os"
	"osLab3/file_sys"
	"strings"
)

func main() {
	fs := file_sys.NewFileSystem()
	currentDir := fs.Root

	pwd := currentDir.Name
	for {
		fmt.Print("[ " + pwd + " ] > ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		// 去除输入中的换行符
		input = strings.TrimSuffix(input, "\n")

		args := strings.Split(input, " ")
		switch args[0] {
		case "mkdir":
			if len(args) != 2 {
				fmt.Println("Usage: mkdir <directory_name>")
				continue
			}
			fs.CreateDirectory(currentDir, args[1], "rwx")

		case "rmdir":
			if len(args) != 2 {
				fmt.Println("Usage: rmdir <directory_name>")
				continue
			}
			fs.DeleteDirectory(currentDir, args[1])

		case "touch":
			if len(args) != 2 {
				fmt.Println("Usage: touch <file_name>")
				continue
			}
			fs.CreateFile(currentDir, args[1], "rw")

		case "rm":
			if len(args) != 2 {
				fmt.Println("Usage: rm <file_name>")
				continue
			}
			fs.DeleteFile(currentDir, args[1])

		case "ls":
			fs.ListDirectory(currentDir)

		case "cd":
			if len(args) != 2 {
				fmt.Println("Usage: cd <dir_name>")
				continue
			}
			if args[1] == ".." {
				if currentDir == fs.Root {
					fmt.Println("已在根目录！")
					continue
				}
				pwd = strings.TrimSuffix(pwd, currentDir.Name+"/")
				currentDir = currentDir.Parent
				continue
			}

			dir, exists := currentDir.SubDirs[args[1]]
			if !exists {
				fmt.Println("目录不存在！")
				continue
			}
			currentDir = dir
			pwd += currentDir.Name + "/"

		case "cat":
			if len(args) != 2 {
				fmt.Println("Usage: cat <file_name>")
				continue
			}
			content := fs.ReadFileContent(currentDir, args[1])
			if content != nil {
				fmt.Println(string(content))
			}

		case "echo":
			if len(args) < 4 || args[len(args)-2] != ">>" {

				for _, a := range args {
					fmt.Println(a)
				}
				fmt.Println("Usage: echo 'content' >> <file_name>")
				continue
			}
			content := strings.Join(args[1:len(args)-2], " ")
			fileName := args[len(args)-1]
			fs.AppendToFile(currentDir, fileName, []byte(content))
		case "tree":
			fs.Tree(currentDir, "")
		default:
			fmt.Println("Unknown command")
		}
	}
}
