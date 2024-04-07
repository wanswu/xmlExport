package cmd

//
//import (
//	"encoding/xml"
//	"fmt"
//	"github.com/xuri/excelize/v2"
//	"io"
//	"os"
//)
//
//type Project struct {
//	GroupID              string       `xml:"groupId"`
//	ArtifactID           string       `xml:"artifactId"`
//	Version              string       `xml:"version"`
//	Dependencies         []Dependency `xml:"dependencies>dependency"`
//	DependencyManagement []Dependency `xml:"dependencyManagement>dependencies>dependency"`
//}
//
//type Dependency struct {
//	GroupID    string `xml:"groupId"`
//	ArtifactID string `xml:"artifactId"`
//	Version    string `xml:"version"`
//}
//
//func ReadXML(filePath string) (project Project) {
//	// 打开并读取 pom.xml 文件
//	file, err := os.Open(filePath)
//	if err != nil {
//		fmt.Println("无法打开 pom.xml 文件:", err)
//		os.Exit(0)
//	}
//	defer file.Close()
//
//	// 读取文件内容
//	content, err := io.ReadAll(file)
//	if err != nil {
//		fmt.Println("无法读取 pom.xml 文件:", err)
//		os.Exit(0)
//	}
//
//	// 解析 XML
//	err = xml.Unmarshal(content, &project)
//	if err != nil {
//		fmt.Println("无法解析 pom.xml 文件:", err)
//		os.Exit(0)
//	}
//	return project
//}
//
//func WriteExecl(project Project) {
//	file := excelize.NewFile()
//	index, err := file.NewSheet("Sheet1")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 设置列名
//	columnNames := []string{"GroupID", "ArtifactID", "Version"}
//	for i, columnName := range columnNames {
//		cell := fmt.Sprintf("%s%d", string('A'+i), 1)
//		file.SetCellValue("Sheet1", cell, columnName)
//	}
//	row := 2 // 从第二行开始写入
//	// 输出 project 的 GroupID、ArtifactID 和 Version
//	file.SetCellValue("Sheet1", "A"+fmt.Sprint(row), project.GroupID)
//	file.SetCellValue("Sheet1", "B"+fmt.Sprint(row), project.ArtifactID)
//	file.SetCellValue("Sheet1", "C"+fmt.Sprint(row), project.Version)
//
//	// 输出 Dependencies.dependencies
//	for _, dependency := range project.Dependencies {
//		row++
//		file.SetCellValue("Sheet1", "A"+fmt.Sprint(row), dependency.GroupID)
//		file.SetCellValue("Sheet1", "B"+fmt.Sprint(row), dependency.ArtifactID)
//		file.SetCellValue("Sheet1", "C"+fmt.Sprint(row), dependency.Version)
//	}
//	for _, dependencys := range project.DependencyManagement {
//		row++
//		file.SetCellValue("Sheet1", "A"+fmt.Sprint(row), dependencys.GroupID)
//		file.SetCellValue("Sheet1", "B"+fmt.Sprint(row), dependencys.ArtifactID)
//		file.SetCellValue("Sheet1", "C"+fmt.Sprint(row), dependencys.Version)
//	}
//	file.SetActiveSheet(index)
//	// 保存 Excel 文件
//	err = file.SaveAs("output.xlsx")
//	if err != nil {
//		fmt.Println("无法保存 Excel 文件:", err)
//		return
//	}
//}
