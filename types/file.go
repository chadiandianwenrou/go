package types

import(
	"os"
	"path"
	"bufio"
	"fmt"
	"github.com/wangbokun/go/log"
	"path/filepath"
	"io/ioutil"
)

// File file
type File struct {
	name  string
	model int
	fd    *os.File
	seek  int64
}


//判断文件目录是否存在
func IsExist(file string) bool{

	_,error	:=	os.Stat(file)

	if error != nil {
		return false
	}
	return true
}

func CreateDir(name string) error {
	return os.MkdirAll(name,os.ModePerm)
}


// create one file
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// remove one file
func Remove(name string) error {
	return os.Remove(name)
}

// get filepath base name
func Basename(fp string) string {
	return path.Base(fp)
}

// get filepath dir name
func Dir(fp string) string {
	return path.Dir(fp)
}

// rename file name
func Rename(src string, target string) error {
	return os.Rename(src, target)
}

// delete file
func Unlink(fp string) error {
	return os.Remove(fp)
}




// list dirs under dirPath
func DirsUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil
}

// list files under dirPath
func FilesUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}

	return ret, nil
}

// get file modified time
func FileMTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = fmt.Errorf("%s not found in paths", fullPath)
	return
}



func ReadLine(r *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := r.ReadLine()
	for isPrefix && err == nil {
		var bs []byte
		bs, isPrefix, err = r.ReadLine()
		line = append(line, bs...)
	}

	return line, err
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
	   fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
	   // 查找文件末尾的偏移量
	   n, _ := f.Seek(0, os.SEEK_END)
	   // 从末尾的偏移量开始写入内容
	   _, err = f.WriteAt([]byte(content), n)
	}  

	defer f.Close()   
	return err

}

// Open open
func (f *File) Open() (err error) {
	f.fd, err = os.OpenFile(f.name, f.model, 0660)
	if err != nil {
		log.Error("<file %s can not open>:%v", f.name, err)
		return
	}
	return
}


func (f *File) Read(p []byte) (int, error) {
	return f.fd.Read(p)
}

func (f *File) Write(p []byte) (int, error) {
	return f.fd.Write(p)
}

// Close close
func (f *File) Close() error {
	return f.fd.Close()
}
