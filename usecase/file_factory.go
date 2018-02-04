package usecase

import (
	"io"
)

type FileFactory interface {
	Create(pre, ext string) io.Writer
	CreateTee(pre, ext string) io.Writer
}

var fileFactory FileFactory

func InitFileFactory(f FileFactory) {
	fileFactory = f
}
