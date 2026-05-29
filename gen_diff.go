package code

import (
	"code/internal/filedata"
)

func GenDiff(path1, path2, format string) (string, error) {
	return filedata.FileData(path1, path2)
}
