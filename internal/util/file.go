package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"git.woa.com/bkdevops/whosbug-ci/internal/zaplog"
)

type FileList []os.FileInfo

func (fileList FileList) Len() int {
	return len(fileList)
}

func (fileList FileList) Swap(i, j int) { // 重写 Swap() 方法
	fileList[i], fileList[j] = fileList[j], fileList[i]
}
func (fileList FileList) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return fileList[j].ModTime().Unix() < fileList[i].ModTime().Unix()
}

func ReadLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	line := ""
	lines = make([]string, 0)
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		line += string(part)
		if !prefix {
			lines = append(lines, line)
			line = ""
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func WriteLines(lines []string, path string) (err error) {
	var file *os.File

	if file, err = os.Create(path); err != nil {
		return
	}

	defer file.Close()

	for _, elem := range lines {
		_, err := file.WriteString(strings.TrimSpace(elem) + "\n")
		if err != nil {
			zaplog.Logger.Error(err.Error())
			break
		}
	}
	return
}

func IsPathExists(path string) (isExists bool, err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		isExists = false
	} else {
		isExists = true
	}
	return
}

func CreateDirIfNotExists(filePath string) (err error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0777)
	}
	return
}

func CreatePathDirIfNotExists(filePath string) (err error) {
	dirPath := path.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0777)
	}
	return
}

func GetFileSize(filePath string) (size int64) {
	size = -1
	fi, e := os.Stat(filePath)
	if e != nil {
		return
	}
	size = fi.Size()
	return
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func WriteBytesToFile(fileContent []byte, path string, fileName string) (filePath string, err error) {

	if err = CreateDirIfNotExists(path); err != nil {
		return
	} else {
		filePath = path + "/" + fileName
		f, err := os.Create(filePath)
		defer f.Close()

		if err != nil {
			return "", err
		} else {
			_, err = f.Write(fileContent)
			return filePath, err
		}

	}

}

func GetBytesFromFile(filePath string) (fileContent []byte, err error) {

	fileContent, err = ioutil.ReadFile(filePath)
	return
}

func GetAllFileWithSuffix(pathname string, suffix string) ([]string, error) {
	var fileList []string
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			if strings.HasSuffix(fi.Name(), suffix) {
				fileList = append(fileList, pathname+"/"+fi.Name())
			}
		}
	}
	return fileList, err
}

func GetAllFileWithSuffixs(pathname string, suffixs []string) ([]string, error) {
	var fileList []string
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			for _, suffix := range suffixs {
				if strings.HasSuffix(fi.Name(), suffix) {
					fileList = append(fileList, pathname+"/"+fi.Name())
				}
			}
		}
	}
	return fileList, err
}

func RemoveFileList(files []string) (isOK bool) {
	filesNum := len(files)
	done := make(chan int, filesNum)
	defer close(done)
	// 这里一定要一个中间变量，不然并发执行的时候，可能删除的都是最后一个文件
	for _, v := range files {
		fileName := v
		go func() {
			if err := os.Remove(fileName); err != nil {
				zaplog.Logger.Error("RemoveFileList error", zaplog.Error(err))
				isOK = false
			}
			done <- 1
		}()
	}
	for i := 0; i < filesNum; i++ {
		<-done
	}
	zaplog.Logger.Info("RemoveFileList finish", zaplog.Int("filesNum", filesNum))
	return true
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteInfoFile(filepath, data string) (err error) {
	fd, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer fd.Close()
	fd.WriteString(data)
	return
}

// GetLineCount 获取文件行数
//
//	@return count
//	@return err
func GetLineCount(path string) (count int64) {
	file, err := os.Open(path)
	if err != nil {
		zaplog.Logger.Error(err.Error())
	}
	defer file.Close()
	fd := bufio.NewReader(file)
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}
	return
}
