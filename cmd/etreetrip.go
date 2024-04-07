package cmd

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

func ReadXMLEtree(file *excelize.File, filePath string, row *int, outFileName string) {
	// 创建excel文件
	index, _ := file.NewSheet("Sheet1")
	file.SetActiveSheet(index)

	// 设置列名
	columnNames := []string{"GroupID", "ArtifactID", "Version"}
	for i, columnName := range columnNames {
		cell := fmt.Sprintf("%s%d", string(rune('A'+i)), 1)
		err := file.SetCellValue("Sheet1", cell, columnName)
		if err != nil {
			return
		}
	}

	// 解析pom文件
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(filePath); err != nil {
		fmt.Println("未找到pom.xml文件")
		os.Exit(0)
	}

	// 查找dependency节点
	for _, dependency := range doc.FindElements("//dependency") {
		groupID := dependency.FindElement("groupId").Text()
		artifactID := dependency.FindElement("artifactId").Text()
		//处理version为空时的情况
		version := "unknown"
		if dependency.FindElement("version") == nil {
		} else {
			version = dependency.FindElement("version").Text()
		}
		//处理版本号为变量的情况
		if strings.Contains(version, "version") {
			for _, temp := range doc.FindElements("//properties/*") {
				if strings.Contains(version, temp.Tag) {
					version = temp.Text()
				}
			}
		}
		err := file.SetCellValue("Sheet1", fmt.Sprintf("A%d", *row), groupID)
		if err != nil {
			fmt.Printf("出错了%s", err)
			os.Exit(0)
		}
		err1 := file.SetCellValue("Sheet1", fmt.Sprintf("B%d", *row), artifactID)
		if err1 != nil {
			fmt.Printf("出错了%s", err1)
			os.Exit(0)
		}
		err2 := file.SetCellValue("Sheet1", fmt.Sprintf("C%d", *row), version)
		if err2 != nil {
			fmt.Printf("出错了%s", err2)
			os.Exit(0)
		}
		*row++
	}

	// 查找plugin节点
	for _, pluginDependency := range doc.FindElements("//plugin") {
		groupID := ""
		if pluginDependency.FindElement("groupId") == nil {
			continue
		} else {
			groupID = pluginDependency.FindElement("groupId").Text()
		}
		artifactID := pluginDependency.FindElement("artifactId").Text()
		//处理version为空时的情况
		version := ""
		if pluginDependency.FindElement("version") == nil {
			version = "unknown"
		} else {
			version = pluginDependency.FindElement("version").Text()
		}
		//处理版本号为变量的情况
		if strings.Contains(version, "version") {
			for _, temp := range doc.FindElements("//properties/*") {
				if strings.Contains(version, temp.Tag) {
					version = temp.Text()
				}
			}
		}
		err := file.SetCellValue("Sheet1", fmt.Sprintf("A%d", *row), groupID)
		if err != nil {
			fmt.Printf("出错了%s", err)
			os.Exit(0)
		}
		err1 := file.SetCellValue("Sheet1", fmt.Sprintf("B%d", *row), artifactID)
		if err1 != nil {
			fmt.Printf("出错了%s", err1)
			os.Exit(0)
		}
		err2 := file.SetCellValue("Sheet1", fmt.Sprintf("C%d", *row), version)
		if err2 != nil {
			fmt.Printf("出错了%s", err2)
			os.Exit(0)
		}
		*row++
	}
	// 保存为excel文件
	if err := file.SaveAs(outFileName); err != nil {
		fmt.Println("Excel文件保存失败:", err)
		os.Exit(0)
	}
}
