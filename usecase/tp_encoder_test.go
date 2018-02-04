package usecase_test

import (
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"github.com/tamura2004/httperf2/usecase"
	"testing"
)

func TestInitTpEncoder(t *testing.T) {
	presenters.InitTpEncoder()

	if usecase.TpEncoder == nil {
		t.Error("bad tp encoder")
	}
}
