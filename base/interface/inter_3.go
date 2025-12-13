package main

import "fmt"

// 基础接口定义
type Reader interface {
	Read() []byte
}

type Writer interface {
	Write(data []byte) error
}

type Closer interface {
	Close() error
}

// 接口组合
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// 文件操作实现
type FileOperation struct {
	filename string
	isOpen   bool
}

func (f *FileOperation) Read() []byte {
	fmt.Printf("读取文件 %s\n", f.filename)
	return []byte("文件内容")
}

func (f *FileOperation) Write(data []byte) error {
	fmt.Printf("向文件 %s 写入数据: %s\n", f.filename, string(data))
	return nil
}

func (f *FileOperation) Close() error {
	f.isOpen = false
	fmt.Printf("关闭文件 %s\n", f.filename)
	return nil
}

// 使用组合接口
func processFile(rwc ReadWriteCloser, data []byte) {
	rwc.Read()
	rwc.Write(data)
	rwc.Close()
}

func main() {
	fileOp := &FileOperation{filename: "test.txt", isOpen: true}
	data := []byte("新的内容")

	processFile(fileOp, data)
}
