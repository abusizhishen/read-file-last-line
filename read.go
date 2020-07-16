package read_file_last_line

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
)

const initReadSize = 2 << 4

func ReadLastLine(fileName string) (data []byte, err error) {
	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		return
	}

	var info os.FileInfo
	info, err = f.Stat()
	if err != nil {
		return
	}

	if info.IsDir() {
		err = fmt.Errorf("invalid file, name:%s", fileName)
		return
	}

	return read(f, info.Size())
}

func read(f *os.File, fileSize int64) (data []byte, err error) {
	var piecesLengthArray []int64
	var buf = bytes.Buffer{}

	var sepIndex int
	var offset, sizeWillRead, sizeHasRead int64
	var b []byte
	var sep = getLineBreak()
	sizeWillRead = initReadSize

	if fileSize < int64(len(sep)) {
		data = make([]byte, fileSize)
		_, err = f.ReadAt(b, 0)
		return
	}

	//ignore the last line break if exists
	b = make([]byte, len(sep))
	if _, err = f.ReadAt(b, fileSize-int64(len(sep))); err != nil {
		return nil, err
	} else if bytes.Equal(b, sep) {
		fileSize -= int64(len(sep))
	}

	for {
		//fmt.Println(piecesLengthArray)
		sizeWillRead = sizeWillRead << 1
		offset = fileSize - (sizeWillRead + sizeHasRead)
		if offset < 0 {
			sizeWillRead = fileSize - sizeHasRead
			offset = 0
		}

		sizeHasRead += sizeWillRead
		b = make([]byte, sizeWillRead)
		_, err = f.ReadAt(b, offset)
		if err != nil {
			return
		}

		//fmt.Println( "length: ",len(b), string(b))
		sepIndex = findSep(b, sep)
		if sepIndex == -1 {
			piecesLengthArray = append(piecesLengthArray, sizeWillRead)
			buf.Write(b)
			if sizeHasRead >= fileSize {
				return bytesReassemble(buf.Bytes(), piecesLengthArray), nil
			}
			continue
		}
		piecesLengthArray = append(piecesLengthArray, int64(len(b)-sepIndex-len(sep)))
		buf.Write(b[sepIndex+len(sep):])
		//fmt.Println(  "sepIndex:", sepIndex)
		//fmt.Println("bytes: ",string(buf.Bytes()))

		return bytesReassemble(buf.Bytes(), piecesLengthArray), nil
	}
}

func findSep(s, sep []byte) int {
	return bytes.LastIndex(s, sep)
}

func getLineBreak() []byte {
	switch runtime.GOOS {
	case "windows":
		return []byte("\r\n")
	default:
		return []byte("\n")
	}
}

// reassemble bytes array, like from [e,d,bc,a] to [a,bc,d,e]
func bytesReassemble(b []byte, piecesLengthArray []int64) (data []byte) {
	var length = int64(len(piecesLengthArray))
	var bytLength = int64(len(b))
	var sum int64
	for _, piecesLength := range piecesLengthArray {
		sum += piecesLength
	}

	if sum != bytLength {
		panic(fmt.Errorf("byes sum is not matched with sum of piecesLengthArray, %d != %d", bytLength, sum))
	}
	data = make([]byte, bytLength)
	var oldStart, start, subLength int64

	for i := length - 1; i >= 0; i-- {
		subLength = piecesLengthArray[i]
		oldStart = bytLength - subLength
		copy(data[start:start+subLength], b[oldStart:oldStart+subLength])
		start += subLength
		bytLength -= subLength
	}

	return data
}