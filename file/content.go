package file

import (
	"io/ioutil"
	"strings"
)

// 文件内容行
func ContentLines(filename string) (lines []string, err error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	lines = strings.Split(string(s), "\n")

	return
}
