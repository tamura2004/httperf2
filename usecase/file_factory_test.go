package usecase_test

import (
	"github.com/tamura2004/httperf2/infra"
	"github.com/tamura2004/httperf2/usecase"
	"testing"
)

func TestInitFileFactory(t *testing.T) {
	infra.InitFileFactory()

	if usecase.FileFactory == nil {
		t.Error("bad init file factory")
	}
}
