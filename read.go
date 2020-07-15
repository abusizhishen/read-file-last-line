package read_file_last_line

import (
	"bytes"
	"fmt"
	"os"
)

const initReadSize = 2<<16
func ReadLastLine(fileName string)(data []byte,err error)  {
	var f *os.File
	f,err = os.Open(fileName)
	if err != nil{
		return
	}

	var info os.FileInfo
	info,err = f.Stat()
	if err != nil{
		return
	}

	if info.IsDir(){
		err = fmt.Errorf("invalid file, name:%s",fileName)
	}

	return 	read(f,info.Size())
}

func read(f *os.File, fileSize int64) (data []byte,err error) {
	var bufIndex []int
	var buf = bytes.Buffer{}

	var i,sepIndex int
	var offset,sizeWillRead, readSize int64
	var b []byte
	var sep = getLineBreak()
	sizeWillRead = initReadSize

	for {
		offset = fileSize-(sizeWillRead+readSize)
		if offset < 0{
			sizeWillRead = fileSize-readSize
			offset = 0
		}

		b = make([]byte, sizeWillRead)
		_,err = f.ReadAt(b,offset)
		if err != nil{
			return
		}

		sepIndex = findSep(b,sep)
		if sepIndex == -1{
			bufIndex = append(bufIndex, int(readSize))
			continue
		}

		b = buf.Bytes()
		data = make([]byte, len(b))
		var length = len(bufIndex)
		var start,end int
		for i=length-1;i>=0;i--{
			copy(data[start:end], b[start:end])
		}

		return
	}
}

func findSep(s, sep []byte) int {
	return bytes.LastIndex(s,sep)
}

func getLineBreak() []byte {
	return []byte("\n")
}