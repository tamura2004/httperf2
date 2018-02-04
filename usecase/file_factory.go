package usecase

import (
	"io"
)

type fileFactory interface {
	Create(pre, ext string) io.Writer
	CreateTee(pre, ext string) io.Writer
}

var FileFactory fileFactory

func InitFileFactory(f fileFactory) {
	FileFactory = f
}
