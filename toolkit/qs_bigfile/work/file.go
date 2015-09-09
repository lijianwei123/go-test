package work

import (
	"os"
	"path"
)

const MAX_FILE_SIZE = 1024 * 1024 * 1024 * 100

type File struct {
	filePath string
}

func NewFile(filePath string) *File {
	return &File{filePath}
}

func (fp *File) GetFilePath() string {
	return fp.filePath
}

//文件大小
func (fp *File) GetFileSize() int64 {
	fileInfo, err := os.Stat(fp.filePath)
	if err != nil {
		panic(err)
	}

	return fileInfo.Size()
}

//文件是否存在
func (fp *File) IsExist() bool {
	_, err := os.Stat(fp.filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//是否文件
func (fp *File) IsFile() bool {
	fileInfo, err := os.Stat(fp.filePath)
	if err == nil && !fileInfo.IsDir() {
		return true
	}
	return false
}

func (fp *File) GetExt() string {
	if fp.IsFile() {
		return path.Ext(fp.filePath)
	}

	return ""
}

type QsFile struct {
	fp *File
}

func NewQsFile(filePath string) *QsFile {
	return &QsFile{&File{filePath}}
}

func (qsFp *QsFile) GetRowNum() int64 {
	return 0
}

type QsChunk struct {
	fpQs     *QsFile
	startPos int64
	endPos   int64
}

func NewQsChunk(fpQs *QsFile, startPos int64, endPos int64) *QsChunk {
	return &QsChunk{fpQs, startPos, endPos}
}

//计算块的分布
func calChunk(fp *QsFile, chunkNum int) (chunks []QsChunk, err error) {
	chunks = []QsChunk{}
	return chunks, nil
}

//搜索
func (chunk *QsChunk) Search(content []string) []string {
	return []string{}
}
