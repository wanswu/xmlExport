package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
	"xmlExport/cmd"
)

func dirPath(temp []os.DirEntry) (xmlPath []string) {
	for _, entry := range temp {
		//if entry.IsDir() {
		//	if !strings.HasPrefix(entry.Name(), ".") {
		//		tempPath, _ := os.Getwd()
		//		nextpath := tempPath + "/" + entry.Name()
		//		fmt.Println(nextpath)
		//	}
		//}
		fileName := entry.Name()

		if strings.HasSuffix(fileName, ".xml") {
			var tempPath string
			tempPath, _ = os.Getwd()
			tempPath = tempPath + "/" + entry.Name()
			xmlPath = append(xmlPath, tempPath)
		}

	}
	return
}
func main() {
	path, _ := os.ReadDir(".")
	outFileName := "output.xlsx"
	row := 2
	fmt.Println("正在读取.....")
	file := excelize.NewFile()
	for _, path := range dirPath(path) {
		cmd.ReadXMLEtree(file, path, &row, outFileName)
		row++
	}
	fmt.Println("Excel保存成功。")
}
