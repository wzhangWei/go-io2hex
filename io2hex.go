package Io2hex

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Hex struct {
	addr   uint32
	length uint32
	data   string
}

//'00' Data Rrecord：用来记录数据，HEX文件的大部分记录都是数据记录
//
//'01' End of File Record:用来标识文件结束，放在文件的最后，标识HEX文件的结尾
//
//'02' Extended Segment Address Record:用来标识扩展段地址的记录
//
//'03' Start Segment Address Record:开始段地址记录
//
//'04' Extended Linear Address Record:用来标识扩展线性地址的记录
//
//'05' Start Linear Address Record:开始线性地址记录

//func Read(filename string) ([]Hex, error){
func Read(filename string) {
	/*获取文件中所以数据*/
	fileData, err := fileRend(filename)
	fmt.Println(fileData, err)

}

func Write(hex []Hex) {

}

func fileRend(fileNme string) (string, error) {
	var fileDataString bytes.Buffer
	fileDates, err := os.Open(fileNme) //os.Open(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer fileDates.Close()

	fileBuf := bufio.NewReader(fileDates)
	if err != nil {
		fmt.Println("Open file error!", fileNme, err)
	}
	for {
		lineData, err := fileBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return "", err
			}
		}
		lineData = strings.TrimSpace(lineData)
		fileDataString.WriteString(lineData)
		fileDataString.WriteString("\n")
	}
	fileDate := fileDataString.String()
	return fileDate, nil
}
