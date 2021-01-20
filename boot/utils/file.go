package utils

import (
	"bufio"
	"github.com/gogf/gf/os/glog"
	"os"
)

func BufferWrite(filename, param string) error {
	fileHandle, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		glog.Error("open file error :", err)
		return err
	}
	defer fileHandle.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriter(fileHandle)
	// 字符串写入
	buf.WriteString(param + "\n")
	// 将缓冲中的数据写入
	err = buf.Flush()
	if err != nil {
		glog.Error("flush error :", err)
		return err
	}
	return nil
}
