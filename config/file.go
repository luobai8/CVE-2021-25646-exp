package config

import (
	"bufio"
	"fmt"
	"os"
)

// filepath 指定文件路径，会返回一个切片类型的数据
func ReadFile(filepath string) []string {
	var FileLines []string
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("文件打开失败：", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		FileLines = append(FileLines, scanner.Text())
	}

	//fmt.Println("文件读取结束")
	return FileLines

}

// filepath 指定要写入的文件路径，value 要写入的内容
func WriteFile(filepath, value string) {

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("写入文件错误：", err)
	}
	defer file.Close()

	file.Write([]byte(value + "\n"))
}
