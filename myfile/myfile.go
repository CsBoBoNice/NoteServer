package myfile

import (
	"fmt"
	"io"
	"os"
)

//函数功能：读取文件
//参数：1，读取的文件名，2，读取到的位置
//返回值：1，读取字节个数，2，是否出错
func ReadFile(name string, buff []byte) (FileByet int64, err error) {
	FileByet = 0             //文件有多少字节
	fi, err := os.Open(name) //打开输入*File 读取文件
	if err != nil {
		panic(err)
	}
	defer fi.Close() //退出后关闭文件
	defer fmt.Printf("读取%s成功\n", name)
	fmt.Printf("正在读取文件\n")
	for {
		n, err := fi.Read(buff) //从input.txt读取
		if err != nil && err != io.EOF {
			panic(err)
		}
		FileByet = FileByet + int64(n)
		if n == 0 {
			break
		}
	}
	fmt.Printf("文件大小=%fMB\n", float64(FileByet)/1024/1024)

	return
}

//函数功能：写入文件
//参数：1，写入的文件名，2，写入的数据 3，写入的字节数
//返回值：1，是否出错
func WriteFile(name string, buff []byte, FileByet int64) (err error) {

	var num int
	fo, err := os.Create(name) //创建输出*File 写文件
	if err != nil {
		panic(err)
	}
	defer fo.Close() //退出后关闭文件

	fmt.Printf("正在写入文件\n")
	num, err = fo.Write(buff[:FileByet])
	if err != nil { //写入output.txt,直到错误 写文件
		panic(err)
	}
	fmt.Printf("写入大小=%fMB\n", float64(num)/1024/1024)

	if int64(num) == FileByet {
		fmt.Printf("写入%s成功\n", name)
	}

	return
}
