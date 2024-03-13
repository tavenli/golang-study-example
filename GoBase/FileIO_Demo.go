package main

import (
	"bufio"
	"fmt"
	"os"
)

func FilesIO_main() {

}

func ReadDirs() {
	workPath, _ := os.Getwd()
	dictDir := fmt.Sprint(workPath, "/")

	entries, err := os.ReadDir(dictDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		info, err := entry.Info()
		fmt.Println(entry.Name(), entry.IsDir())
		fmt.Println(info, err)

	}
}

func ReadDict(filePath string) {

	// 打开文件并创建一个新的scanner对象
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func WriteDict(filePath string, datas []string) {

	fWrite, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fWrite.Close()

	for _, v := range datas {
		fWrite.WriteString(fmt.Sprintln(v))
	}

}
