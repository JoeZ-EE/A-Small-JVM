package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, error := filepath.Abs(path)
	if error != nil {
		panic(error)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, error := ioutil.ReadFile(fileName)
	return data, self, error
}

func (self *DirEntry) String() string {
	return self.absDir
}