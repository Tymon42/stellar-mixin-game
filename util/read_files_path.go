package util

import (
	"io/ioutil"
	"os"
	"strings"
)

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth, filetpye string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name(), filetpye)
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), filetpye)
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table,filetpye)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}