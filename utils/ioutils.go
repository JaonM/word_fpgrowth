// 文件读写、json读写、数据库读写等待
package utils

import (
	"bufio"
	"encoding/csv"
	"os"
)

func ReadFile(filePath string) [] string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var texts [] string
	for sc.Scan() {
		texts = append(texts, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return texts
}

func WriteCSV(filePath string, headers [] string, data [][] string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	if headers != nil {
		err := writer.Write(headers)
		if err != nil {
			panic(err)
		}
	}
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			panic(err)
		}
	}
}

