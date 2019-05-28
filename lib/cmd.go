package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ViewFiles(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	names := make([]string, 0)
	for _, f := range files {
		names = append(names, f.Name())
	}
	return names
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func OutputContent(filename string) string {
	// 读文件
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
