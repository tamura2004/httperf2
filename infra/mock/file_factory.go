package mock

import (
	"github.com/tamura2004/httperf2/infra"
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"io/ioutil"
)

func CreateFile(pre, ext string) io.Writer {
	return ioutil.Discard
}

func InitFileFactory() {
	usecase.FileFactory = infra.FileFactory(CreateFile)
}
