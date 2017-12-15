package myfile

import (
	"fmt"
	"io"
	"os"
)

//使用的例子
func UsageMethod() {
	const DateSize = 1024 * 1024 * 32 //每次处理的数据大小

	input := "E:/input.mp4"
	output := "E:/output.mp4"

	var FileSize int64
	var Date []byte = make([]byte, DateSize) //2G内存

	fi, err := os.Open(input) //打开输入*File 读取文件
	CheckFile(err)
	defer fi.Close() //退出后关闭文件

	fo, err := os.Create(output) //创建输出*File 写文件
	CheckFile(err)
	defer fo.Close() //退出后关闭文件

	FileSize = GetFileSize(fi)
	fmt.Printf("文件大小=%fMB\n", float64(FileSize)/1024/1024)

	fornum := FileSize / DateSize
	if (FileSize % DateSize) > 0 {
		fornum++
	}
	OverplusDate := FileSize - (DateSize * (fornum - 1))
	for i := 0; i < int(fornum); i++ {
		_, err := fi.Seek(0, 0)
		CheckFile(err)
		_, err = fo.Seek(0, 0)
		CheckFile(err)

		if i != int(fornum)-1 {
			err = ReadPart(fi, int64(DateSize*i), Date, DateSize)
			CheckFile(err)
			err = WritePart(fo, int64(DateSize*i), Date, DateSize)
			CheckFile(err)
		} else {
			err = ReadPart(fi, int64(DateSize*i), Date, OverplusDate)
			CheckFile(err)
			err = WritePart(fo, int64(DateSize*i), Date, OverplusDate)
			CheckFile(err)
		}
	}
	if FileSize == GetFileSize(fo) {
		fmt.Printf("读取数据与写入数据相同~nice code\n")
	}
}

//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。
func CheckFile(e error) {
	if e != nil {
		panic(e)
	}
}

//得到文件的字节大小
//返回文件字节大小
func GetFileSize(fd *os.File) (size int64) {

	this, err := fd.Seek(0, 1) //保存当前位置
	CheckFile(err)
	_, err = fd.Seek(0, 0) //指向文件头
	CheckFile(err)
	size, err = fd.Seek(0, 2) //得到文件字节大小
	CheckFile(err)
	_, err = fd.Seek(this, 0) //回到原来文件指向位置
	CheckFile(err)

	return
}

//函数功能：读取部分文件
//参数：1，读取的文件指针，2，读取的偏移量 3，存取的位置 4，读取的字节个数
//返回值：1，是否出错
func ReadPart(fd *os.File, ret int64, buff []byte, Size int64) (err error) {
	var FileByet int64
	_, err = fd.Seek(0, 0)
	CheckFile(err)
	_, err = fd.Seek(ret, 0)
	CheckFile(err)
	n, err := fd.Read(buff) //从input.txt读取
	if err != nil && err != io.EOF {
		if int64(n) != Size {
			panic(err)
		}
	}
	FileByet = FileByet + int64(n)
	fmt.Printf("读取大小=%fMB\n", float64(FileByet)/1024/1024)
	if FileByet == int64(Size) {
		err = nil
		fmt.Printf("成功读取%fMB\n", float64(FileByet)/1024/1024)
		return
	}
	err = fmt.Errorf("%s", "read FileByet!=size")
	return
}

//函数功能：写入部分文件
//参数：1，写入的文件指针，2，写入的偏移量 3，写入的位置 4，写入的字节个数
//返回值：1，是否出错
func WritePart(fd *os.File, ret int64, buff []byte, Size int64) (err error) {

	_, err = fd.Seek(0, 0)
	CheckFile(err)
	_, err = fd.Seek(ret, 0)
	CheckFile(err)
	n, err := fd.Write(buff[:Size])
	if err != nil { //写入output.txt,直到错误 写文件
		panic(err)
	}
	fmt.Printf("写入大小=%fMB\n", float64(n)/1024/1024)

	if int64(n) == Size {
		err = nil
		fmt.Printf("写入成功\n")
		return
	}
	err = fmt.Errorf("%s", "Write n!=size")
	return
}

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
