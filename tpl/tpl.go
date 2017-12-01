package tpl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"text/template"
	"os"
	"path/filepath"
	"fmt"
)

func Handler(data interface{}, tplFile, outFile string) error {
	var err error
	fmt.Println("---------outFile----------", outFile)
	tplData, err := ioutil.ReadFile(tplFile)
	if err != nil {
		return err
	}

	ret, _ := json.Marshal(data)
	data = nil
	json.Unmarshal(ret, &data)

	buf := bytes.NewBuffer(nil)
	t := template.Must(template.New("data").Funcs(template.FuncMap{
		"add": add,
		"len": sliceLen,
	}).Parse(string(tplData)))

	t.Execute(buf, data)

	newBuf, err := stripNullLine(buf)
	if err != nil {
		return err
	}

	err = WriteFile(outFile, newBuf)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(fileName string, data []byte) error {
	//创建多级目录文件
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if err := CreatFile(fileName); err != nil {
				return err
			}
			return ioutil.WriteFile(fileName, data, 0755)
		} else if os.IsExist(err) {
			return ioutil.WriteFile(fileName, data, 0755)
		} else {
			return nil
		}
	}
	return ioutil.WriteFile(fileName, data, 0755)
}

func CreatFile(fileName string) error {
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			// Directory does not exist, create it //注意只能创建目录 filePath.Dir
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			if _, err := os.Create(fileName); err != nil {
				return err
			}
		} else if os.IsExist(err) {
			if _, err := os.Create(fileName); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// self define functions
func add(value, addValue int) int {
	return value + addValue
}

func sliceLen(slice []interface{}) int {
	return len(slice)
}

// strip null lines
func stripNullLine(buf *bytes.Buffer) ([]byte, error) {
	newBuf := bytes.NewBuffer(nil)

	for {
		str, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		if len(str) <= 1 || str == "\r\n" {
			continue
		}
		newBuf.WriteString(str)
	}
	return newBuf.Bytes(), nil
}
