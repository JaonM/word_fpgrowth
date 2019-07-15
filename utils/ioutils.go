// 文件读写、json读写、数据库读写等待
package utils

import (
	"bufio"
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
